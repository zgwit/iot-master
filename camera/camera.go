package camera

import (
	"context"
	"fmt"
	"github.com/zgwit/iot-master/lib"
	"github.com/zgwit/iot-master/model"
	"os"
	"os/exec"
)

type Segment struct {
	name string
	data []byte
}

type Camera struct {
	model.Camera

	running bool

	cancel context.CancelFunc

	playlist []byte
	segments lib.LinkList[*Segment]
}

func NewCamera(m *model.Camera) *Camera {
	return &Camera{
		Camera: *m,
	}
}

func (c *Camera) Running() bool {
	return c.running
}

func (c *Camera) Open() error {
	if c.running {
		_ = c.Close()
	}

	var ctx context.Context
	ctx, c.cancel = context.WithCancel(context.Background())
	//TODO 转码，压缩，省流量
	cmd := exec.CommandContext(ctx,
		"ffmpeg", "-re", "-i", c.Url,
		//"-c:a", "aac",
		"-c:v", "libx264", //使用H264，因为前端仅支持h264
		//"-r", "10",
		"-f", "hls",
		"-hls_allow_cache", "0",
		//"-hls_flags", "delete_segments", //windows下 路径分隔符错误，导致不能删除，另外可能不支持http删除，需要手动操作
		fmt.Sprintf("http://localhost:143/%d/index.m3u8", c.Id),
	)
	//cmd = exec.Command("ffmpeg", "-re", "-i", c.MediaUri, "-c", "copy", "-f", "hls", "http://localhost:8080/camera/index.m3u8")
	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	//cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return err
	}
	go func() {
		c.running = true
		_ = cmd.Wait()
		c.running = false
	}()

	return nil
}

func (c *Camera) Playlist() []byte {
	return c.playlist
}

func (c *Camera) Segment(name string) []byte {
	var data []byte
	c.segments.Walk(func(segment *Segment) bool {
		if name == segment.name {
			data = segment.data
			return false
		}
		return true
	})
	return data
}

func (c *Camera) Close() error {
	if c.cancel != nil {
		c.cancel()
		c.cancel = nil
	}
	return nil
}
