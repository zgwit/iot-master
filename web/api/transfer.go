package api

import (
	"github.com/gin-gonic/gin"
	"iot-master/db"
	"iot-master/master"
	"iot-master/model"
)

func transferList(ctx *gin.Context) {
	transfers := make([]*model.TransferEx, 0)

	var body paramSearchEx
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	query := body.toQuery()
	query.Select("transfer.*, " + //TODO 只返回需要的字段
		" 0 as running, tunnel.name as tunnel")
	query.Join("LEFT", "tunnel", "transfer.tunnel_id=tunnel.id")

	cnt, err := query.FindAndCount(&transfers)
	if err != nil {
		replyError(ctx, err)
		return
	}
	for _, pe := range transfers {
		d := master.GetTransfer(pe.Id)
		if d != nil {
			pe.Running = d.Running()
		}
	}
	replyList(ctx, transfers, cnt)
}

func afterTransferCreate(data interface{}) error {
	transfer := data.(*model.Transfer)
	//启动
	if !transfer.Disabled {
		return master.LoadTransfer(transfer.Id)
	}
	return nil
}

func transferDetail(ctx *gin.Context) {
	var te model.TransferEx
	has, err := db.Engine.ID(ctx.GetInt64("id")).Get(&te.Transfer)
	if err != nil {
		replyError(ctx, err)
		return
	}
	if !has {
		replyFail(ctx, "记录不存在")
		return
	}
	d := master.GetTransfer(te.Id)
	if d != nil {
		te.Running = d.Running()
	}
	replyOk(ctx, te)

}

func afterPipeUpdate(data interface{}) error {
	transfer := data.(*model.Transfer)
	_ = master.RemoveTransfer(transfer.Id)
	return master.LoadTransfer(transfer.Id)
}

func afterTransferDelete(id interface{}) error {
	return master.RemoveTransfer(id.(int64))
}

func transferStart(ctx *gin.Context) {
	transfer := master.GetTransfer(ctx.GetInt64("id"))
	if transfer == nil {
		replyFail(ctx, "not found")
		return
	}
	err := transfer.Open()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func transferStop(ctx *gin.Context) {
	transfer := master.GetTransfer(ctx.GetInt64("id"))
	if transfer == nil {
		replyFail(ctx, "not found")
		return
	}
	err := transfer.Close()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func afterTransferEnable(id interface{}) error {
	_ = master.RemoveTransfer(id.(int64))
	return master.LoadTransfer(id.(int64))
}

func afterTransferDisable(id interface{}) error {
	return master.RemoveTransfer(id.(int64))
}
