package camera

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/aler9/gortsplib"
	"github.com/aler9/gortsplib/pkg/h264"
	"github.com/grafov/m3u8"
	"github.com/zgwit/gonvif"
	"github.com/zgwit/gonvif/media"
	"github.com/zgwit/iot-master/model"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

type Camera struct {
	model.Camera

	playlist  m3u8.Playlist
	client    *gortsplib.Client
	extractor *h264.DTSExtractor
	segment   *Segment
	last      time.Duration
	startPCR  time.Time
	startDTS  time.Duration
}

func (c *Camera) getMediaUri() error {
	dev, err := gonvif.NewDevice(
		gonvif.DeviceParams{
			Xaddr:      c.Address,
			Username:   c.Username,
			Password:   c.Password,
			HttpClient: nil,
		})
	if err != nil {
		return err
	}
	resp, err := dev.CallMethod(media.GetStreamUri{ProfileToken: "Profile_101"})
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var msu media.GetStreamUriResponse
	err = xml.Unmarshal(b, &msu)
	if err != nil {
		return err
	}

	c.MediaUri = string(msu.MediaUri.Uri)
	return nil
}

func (c Camera) PrimaryPlaylist() string {

	var codecs []string

	//gortsplib.NewTrackH264()
	if p.videoTrack != nil {
		sps := p.videoTrack.SPS()
		if len(sps) >= 4 {
			codecs = append(codecs, "avc1."+hex.EncodeToString(sps[1:4]))
		}
	}

	list := "#EXTM3U\n" +
		"#EXT-X-VERSION:3\n" +
		"#EXT-X-INDEPENDENT-SEGMENTS\n" +
		"\n" +
		"#EXT-X-STREAM-INF:BANDWIDTH=200000,CODECS=\"" + strings.Join(codecs, ",") + "\"\n" +
		"stream.m3u8\n"
	return list
}

func (c *Camera) Playlist() string {
	list := "#EXTM3U\n"
	list += "#EXT-X-VERSION:3\n"
	list += "#EXT-X-ALLOW-CACHE:NO\n"

	targetDuration := func() uint {
		ret := uint(0)

		// EXTINF, when rounded to the nearest integer, must be <= EXT-X-TARGETDURATION
		for _, s := range c.segments {
			v2 := uint(math.Round(s.duration().Seconds()))
			if v2 > ret {
				ret = v2
			}
		}

		return ret
	}()
	list += "#EXT-X-TARGETDURATION:" + strconv.FormatUint(uint64(targetDuration), 10) + "\n"

	list += "#EXT-X-MEDIA-SEQUENCE:" + strconv.FormatInt(int64(p.segmentDeleteCount), 10) + "\n"
	list += "\n"

	for _, s := range p.segments {
		list += "#EXT-X-PROGRAM-DATE-TIME:" + s.startTime.Format("2006-01-02T15:04:05.999Z07:00") + "\n" +
			"#EXTINF:" + strconv.FormatFloat(s.duration().Seconds(), 'f', -1, 64) + ",\n" +
			s.name + ".ts\n"
	}

	return list
}

func (c *Camera) OpenStream() error {
	c.client = &gortsplib.Client{
		OnPacketRTP: func(ctx *gortsplib.ClientOnPacketRTPCtx) {

			if ctx.H264NALUs == nil {
				return
			}

			idr := h264.IDRPresent(ctx.H264NALUs)

			var err error
			var dts, pts time.Duration
			if c.segment == nil {
				if !idr {
					return
				}

				c.extractor = h264.NewDTSExtractor()
				dts, err = c.extractor.Extract(ctx.H264NALUs, ctx.H264PTS)
				if err != nil {
					fmt.Println(err)
				}

				c.startPCR = time.Now()
				c.startDTS = dts

				dts = 0
				pts = ctx.H264PTS - c.startDTS

				c.segment = newSegment(nil, nil)

			} else {
				dts, err = c.extractor.Extract(ctx.H264NALUs, ctx.H264PTS)
				if err != nil {
					fmt.Println(err)
				}

				dts -= c.startDTS
				pts = ctx.H264PTS - c.startDTS

				if idr && dts-c.last > 10*time.Second {
					//ready
					c.segment = newSegment(nil, nil)
				}
			}

			err = c.segment.writeH264(
				time.Now().Sub(c.startPCR),
				dts,
				pts, idr, ctx.H264NALUs)
			if err != nil {
				return
			}

			//c.segment.buf
		},
	}

	return c.client.StartReading(c.MediaUri)
}

func (c *Camera) Wait() error {
	return c.client.Wait()
}

func (c *Camera) Start() error {

	return nil
}
