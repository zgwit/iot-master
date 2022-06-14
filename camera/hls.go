package camera

import (
	"bufio"
	"context"
	"github.com/aler9/gortsplib"
	"github.com/aler9/gortsplib/pkg/h264"
	"github.com/asticode/go-astits"
)

func te() {
	c := &gortsplib.Client{
		OnPacketRTP: func(ctx *gortsplib.ClientOnPacketRTPCtx) {
			//ctx.H264NALUs
			//ctx.H264PTS
			idr := h264.IDRPresent(ctx.H264NALUs)
			dts := h264.NewDTSExtractor()
			dts.Extract(ctx.H264NALUs, ctx.H264PTS)

			nalus := append([][]byte{{byte(h264.NALUTypeAccessUnitDelimiter), 240}}, ctx.H264NALUs...)
			enc, _ := h264.AnnexBEncode(nalus)

			af := &astits.PacketAdaptationField{RandomAccessIndicator: idr}

			oh := &astits.PESOptionalHeader{
				MarkerBits: 2,
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
}
