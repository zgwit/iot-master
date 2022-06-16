package camera

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/aler9/gortsplib"
	"github.com/aler9/gortsplib/pkg/h264"
	"github.com/asticode/go-astits"
)

const (
	mpegtsPCROffset = 400 * time.Millisecond // 2 samples @ 5fps
)

type Segment struct {
	videoTrack     *gortsplib.TrackH264
	writeData      func(*astits.MuxerData) (int, error)
	buf            bytes.Buffer
	pcrSendCounter int
}

func newSegment(videoTrack *gortsplib.TrackH264, writeData func(*astits.MuxerData) (int, error)) *Segment {
	t := &Segment{
		videoTrack: videoTrack,
		writeData:  writeData,
	}

	// WriteTable() is called automatically when WriteData() is called with
	// - PID == PCRPID
	// - AdaptationField != nil
	// - RandomAccessIndicator = true

	return t
}

func (t *Segment) write(p []byte) (int, error) {
	if uint64(len(p)+t.buf.Len()) > t.segmentMaxSize {
		return 0, fmt.Errorf("reached maximum segment size")
	}

	return t.buf.Write(p)
}

func (t *Segment) reader() io.Reader {
	return bytes.NewReader(t.buf.Bytes())
}

func (t *Segment) writeH264(pcr, dts, pts time.Duration, idrPresent bool, nalus [][]byte) error {
	// prepend an AUD. This is required by video.js and iOS
	nalus = append([][]byte{{byte(h264.NALUTypeAccessUnitDelimiter), 240}}, nalus...)

	enc, err := h264.AnnexBEncode(nalus)
	if err != nil {
		return err
	}

	var af *astits.PacketAdaptationField

	if idrPresent {
		af = &astits.PacketAdaptationField{}
		af.RandomAccessIndicator = true
	}

	// send PCR once in a while
	if t.pcrSendCounter == 0 {
		if af == nil {
			af = &astits.PacketAdaptationField{}
		}
		af.HasPCR = true
		af.PCR = &astits.ClockReference{Base: int64(pcr.Seconds() * 90000)}
		t.pcrSendCounter = 3
	}
	t.pcrSendCounter--

	oh := &astits.PESOptionalHeader{
		MarkerBits: 2,
	}

	if dts == pts {
		oh.PTSDTSIndicator = astits.PTSDTSIndicatorOnlyPTS
		oh.PTS = &astits.ClockReference{Base: int64((pts + mpegtsPCROffset).Seconds() * 90000)}
	} else {
		oh.PTSDTSIndicator = astits.PTSDTSIndicatorBothPresent
		oh.DTS = &astits.ClockReference{Base: int64((dts + mpegtsPCROffset).Seconds() * 90000)}
		oh.PTS = &astits.ClockReference{Base: int64((pts + mpegtsPCROffset).Seconds() * 90000)}
	}

	_, err = t.writeData(&astits.MuxerData{
		PID:             256,
		AdaptationField: af,
		PES: &astits.PESData{
			Header: &astits.PESHeader{
				OptionalHeader: oh,
				StreamID:       224, // video
			},
			Data: enc,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
