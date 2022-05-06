package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/log"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
)

func pipeRoutes(app *gin.RouterGroup) {
	app.POST("list", pipeList)
	app.POST("create", pipeCreate)

	app.Use(parseParamId)
	app.GET(":id", pipeDetail)
	app.POST(":id", pipeUpdate)
	app.GET(":id/delete", pipeDelete)
	app.GET(":id/start", pipeStart)
	app.GET(":id/stop", pipeStop)
	app.GET(":id/enable", pipeEnable)
	app.GET(":id/disable", pipeDisable)
}

func pipeList(ctx *gin.Context) {
	records, cnt, err := normalSearch(ctx, database.Master, &model.Pipe{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	//补充信息
	pipes := records.(*[]*model.Pipe)
	ps := make([]*model.PipeEx, 0) //len(pipes)

	for _, d := range *pipes {
		pe := &model.PipeEx{Pipe: *d}
		ps = append(ps, pe)
		p := master.GetPipe(pe.Id)
		if p != nil {
			pe.Running = p.Running()
		}
		var link model.Link
		err = database.Master.One("Id", d.LinkId, &link)
		if err == nil {
			if link.Name != "" {
				pe.Link = link.Name
			} else if link.SN != "" {
				pe.Link = link.SN
			} else if link.Remote != "" {
				pe.Link = link.Remote
			}
		}
	}

	replyList(ctx, ps, cnt)
}

func pipeCreate(ctx *gin.Context) {
	var pipe model.Pipe
	err := ctx.ShouldBindJSON(&pipe)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = database.Master.Save(&pipe)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, pipe)

	//启动
	//pipeStart(ctx)
	if !pipe.Disabled {
		go func() {
			err := master.LoadPipe(pipe.Id)
			if err != nil {
				log.Error(err)
				return
			}
		}()
	}
}

func pipeDetail(ctx *gin.Context) {
	var pipe model.Pipe
	err := database.Master.One("Id", ctx.GetInt("id"), &pipe)
	if err != nil {
		replyError(ctx, err)
		return
	}
	pe := &model.PipeEx{Pipe: pipe}
	d := master.GetPipe(pe.Id)
	if d != nil {
		pe.Running = d.Running()
	}
	var link model.Link
	err = database.Master.One("Id", pipe.LinkId, &link)
	if err == nil {
		if link.Name != "" {
			pe.Link = link.Name
		} else if link.SN != "" {
			pe.Link = link.SN
		} else if link.Remote != "" {
			pe.Link = link.Remote
		}
	}
	replyOk(ctx, pe)
}

func pipeUpdate(ctx *gin.Context) {
	var pipe model.Pipe
	err := ctx.ShouldBindJSON(&pipe)
	if err != nil {
		replyError(ctx, err)
		return
	}
	pipe.Id = ctx.GetInt("id")

	err = database.Master.Update(&pipe)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, pipe)

	//重新启动
	go func() {
		_ = master.RemovePipe(pipe.Id)
		err := master.LoadPipe(pipe.Id)
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func pipeDelete(ctx *gin.Context) {
	pipe := model.Pipe{Id: ctx.GetInt("id")}
	err := database.Master.DeleteStruct(&pipe)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, pipe)

	//关闭
	go func() {
		err := master.RemovePipe(pipe.Id)
		if err != nil {
			log.Error(err)
		}
	}()
}

func pipeStart(ctx *gin.Context) {
	pipe := master.GetPipe(ctx.GetInt("id"))
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
	pipe := master.GetPipe(ctx.GetInt("id"))
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

func pipeEnable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Pipe{Id: ctx.GetInt("id")}, "Disabled", false)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

	//启动
	go func() {
		err := master.LoadPipe(ctx.GetInt("id"))
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func pipeDisable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Pipe{Id: ctx.GetInt("id")}, "Disabled", true)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

	//关闭
	go func() {
		pipe := master.GetPipe(ctx.GetInt("id"))
		if pipe == nil {
			return
		}
		err := pipe.Close()
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

