package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/storm/v3/q"
)

func userRoutes(app *gin.RouterGroup) {
	app.POST("list", userList)
	app.POST("create", userCreate)

	app.GET("event/clear", userEventClearAll)

	app.Use(parseParamId)

	app.GET(":id", userDetail)
	app.POST(":id", userUpdate)
	app.GET(":id/delete", userDelete)
	app.GET(":id/password", userPassword)
	app.GET(":id/enable", userEnable)
	app.GET(":id/disable", userDisable)

	app.POST(":id/event/list", userEvent)
	app.GET(":id/event/clear", userEventClear)
}

func userList(ctx *gin.Context) {
	users, cnt, err := normalSearch(ctx, database.Master, &model.User{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, users, cnt)
}

func userCreate(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = database.Master.Save(&user)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 默认密码

	replyOk(ctx, user)
}


func userDetail(ctx *gin.Context) {
	var user model.User
	err := database.Master.One("Id", ctx.GetInt("id"), &user)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, user)
}

func userUpdate(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		replyError(ctx, err)
		return
	}
	user.Id = ctx.GetInt("id")

	err = database.Master.Update(&user)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, user)
}

func userDelete(ctx *gin.Context) {
	user := model.User{Id: ctx.GetInt("id")}
	err := database.Master.DeleteStruct(&user)
	if err != nil {
		replyError(ctx, err)
		return
	}

	_= database.Master.DeleteStruct(model.Password{Id: user.Id})

	replyOk(ctx, user)
}

func userPassword(ctx *gin.Context) {

	replyOk(ctx, nil)
}

func userEnable(ctx *gin.Context) {
	err := database.Master.UpdateField(model.User{Id: ctx.GetInt("id")}, "Disabled", false)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}

func userDisable(ctx *gin.Context) {
	err := database.Master.UpdateField(model.User{Id: ctx.GetInt("id")}, "Disabled", true)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}

func userEvent(ctx *gin.Context) {
	events, cnt, err := normalSearchById(ctx, database.History, "UserId", ctx.GetInt("id"), &model.UserEvent{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, events, cnt)
}

func userEventClear(ctx *gin.Context) {
	err := database.History.Select(q.Eq("UserId", ctx.GetInt("id"))).Delete(&model.UserEvent{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func userEventClearAll(ctx *gin.Context) {
	err := database.History.Drop(&model.UserEvent{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}
