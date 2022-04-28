package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/storm/v3"
	"github.com/zgwit/storm/v3/q"
	"time"
)

func visitRoutes(app *gin.RouterGroup) {
	app.POST("list", visitList)
	app.POST("create", visitCreate)
	app.GET("clear", visitClear)
	app.GET("clear/all", visitClearAll)

}


func visitList(ctx *gin.Context) {
	history, cnt, err := normalSearch(ctx, database.History, &model.Visit{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, history, cnt)
}

func visitCreate(ctx *gin.Context) {
	var history, last model.Visit
	err := ctx.ShouldBindJSON(&history)
	if err != nil {
		replyError(ctx, err)
		return
	}

	user := ctx.MustGet("user").(*model.User)

	err = database.Master.Select(q.Eq("UserId", user.Id), q.Eq("TargetId", history.TargetId), q.Eq("Target", history.Target)).First(&last)
	if err == storm.ErrNotFound {
		err = database.Master.Save(&history)
	} else if err == nil {
		err = database.Master.UpdateField(&last, "Last", time.Now())
	}

	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func visitClear(ctx *gin.Context) {
	user := ctx.MustGet("user").(*model.User)
	err := database.History.Select(q.Eq("UserId", user.Id)).Delete(&model.Visit{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func visitClearAll(ctx *gin.Context) {
	err := database.History.Drop(&model.Visit{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}
