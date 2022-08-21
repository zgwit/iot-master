package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/internal/core"
	"github.com/zgwit/iot-master/model"
)

func afterServerCreate(data interface{}) error {
	server := data.(*model.Server)
	if !server.Disabled {
		return core.LoadServer(server.Id)
	}
	return nil
}

func afterServerUpdate(data interface{}) error {
	server := data.(*model.Server)
	_ = core.RemoveServer(server.Id)
	return core.LoadServer(server.Id)
}

func afterServerDelete(id interface{}) error {
	return core.RemoveServer(id.(uint64))
}

func serverStart(ctx *gin.Context) {
	server := core.GetServer(ctx.GetUint64("id"))
	if server == nil {
		replyFail(ctx, "not found")
		return
	}
	err := server.Instance.Open()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func serverStop(ctx *gin.Context) {
	server := core.GetServer(ctx.GetUint64("id"))
	if server == nil {
		replyFail(ctx, "not found")
		return
	}
	err := server.Instance.Close()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func afterServerEnable(id interface{}) error {
	_ = core.RemoveServer(id.(uint64))
	return core.LoadServer(id.(uint64))
}

func afterServerDisable(id interface{}) error {
	return core.RemoveServer(id.(uint64))
}
