package camera

import (
	"bufio"
	"context"
	"github.com/aler9/gortsplib"
	"github.com/aler9/gortsplib/pkg/h264"
	"github.com/asticode/go-astits"
	"time"
)

const (
//mpegtsPCROffset = 400 * time.Millisecond // 2 samples @ 5fps
)

func TestRTSP() {
	c := &gortsplib.Client{
		OnPacketRTP: func(ctx *gortsplib.ClientOnPacketRTPCtx) {

			//ctx.H264NALUs
			//ctx.H264PTS
			idr := h264.IDRPresent(ctx.H264NALUs)
			x := h264.NewDTSExtractor()

			_, err := x.Extract(ctx.H264NALUs, ctx.H264PTS)
			if err != nil {
				return
			}

			var pcr, dts, pts time.Duration

			// prepend an AUD. This is required by video.js and iOS
			nalus := append([][]byte{{byte(h264.NALUTypeAccessUnitDelimiter), 240}}, ctx.H264NALUs...)
			enc, err := h264.AnnexBEncode(nalus)

			af := &astits.PacketAdaptationField{RandomAccessIndicator: idr}
			if true {
				af.HasPCR = true
				af.PCR = &astits.ClockReference{Base: int64(pcr.Seconds() * 90000)}
			}

			oh := &astits.PESOptionalHeader{
				MarkerBits: 2,
				PTS:        &astits.ClockReference{Base: int64((pts + mpegtsPCROffset).Seconds() * 90000)},
			}

			if dts == pts {
				oh.PTSDTSIndicator = astits.PTSDTSIndicatorOnlyPTS
			} else {
				oh.PTSDTSIndicator = astits.PTSDTSIndicatorBothPresent
				oh.DTS = &astits.ClockReference{Base: int64((dts + mpegtsPCROffset).Seconds() * 90000)}
			}

			md := &astits.MuxerData{
				PID:             256,
				AdaptationField: af,
				PES: &astits.PESData{
					Header: &astits.PESHeader{
						OptionalHeader: oh,
						StreamID:       224, // video
					},
					Data: enc,
				},
			}

			m := astits.NewMuxer(context.Background(), bufio.NewWriter(nil))
			_, _ = m.WriteData(md)
		},
	}

	err := c.StartReading("rtsp://localhost:8554/mystream")
	if err != nil {
		return
	}

	c.Wait()
}
