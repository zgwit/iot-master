package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
)

func pipeList(ctx *gin.Context) {
	pipes := make([]*model.PipeEx, 0)

	var body paramSearchEx
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	query := body.toQuery()

	query.Join("LEFT", "link", "pipe.link_id=link.id")

	cnt, err := query.FindAndCount(pipes)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, pipes, cnt)
}

func afterPipeCreate(data interface{}) error {
	pipe := data.(*model.Pipe)
	//启动
	if !pipe.Disabled {
		return master.LoadPipe(pipe.Id)
	}
	return nil
}

func pipeDetail(ctx *gin.Context) {
	var pe model.PipeEx
	var link model.Link
	pe.Link = link.Name
	pe.Link = link.SN
	pe.Link = link.Remote

	replyOk(ctx, pe)
}

func afterPipeUpdate(data interface{}) error {
	pipe := data.(*model.Pipe)
	_ = master.RemovePipe(pipe.Id)
	return master.LoadPipe(pipe.Id)
}

func afterPipeDelete(data interface{}) error{
	pipe := data.(*model.Pipe)
	return master.RemovePipe(pipe.Id)
}

func pipeStart(ctx *gin.Context) {
	pipe := master.GetPipe(ctx.GetInt64("id"))
	if pipe == nil {
		replyFail(ctx, "not found")
		return
	}
	err := pipe.Open()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func pipeStop(ctx *gin.Context) {
	pipe := master.GetPipe(ctx.GetInt64("id"))
	if pipe == nil {
		replyFail(ctx, "not found")
		return
	}
	err := pipe.Close()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func afterPipeEnable(data interface{}) error{
	pipe := data.(*model.Pipe)
	_ = master.RemovePipe(pipe.Id)
	return master.LoadPipe(pipe.Id)
}

func afterPipeDisable(data interface{}) error{
	pipe := data.(*model.Pipe)
	return master.RemovePipe(pipe.Id)
}
