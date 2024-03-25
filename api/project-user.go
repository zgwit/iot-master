package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v4/pkg/db"
	"github.com/zgwit/iot-master/v4/project"
	"github.com/zgwit/iot-master/v4/web/curd"
	"xorm.io/xorm/schemas"
)

// @Summary 项目用户列表
// @Schemes
// @Description 项目用户列表
// @Tags project-user
// @Param id path int true "项目ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[[]project.ProjectUser] 返回项目用户信息
// @Router /project/{id}/user/list [get]
func projectUserList(ctx *gin.Context) {
	var pds []project.ProjectUser
	err := db.Engine.
		Select("project_user.project_id, project_user.user_id, project_user.admin, project_user.disabled, project_user.created, user.name as user").
		Join("INNER", "user", "user.id=project_user.user_id").
		Where("project_user.project_id=?", ctx.Param("id")).
		Find(&pds)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, pds)
}

// @Summary 判断项目用户是否存在
// @Schemes
// @Description 判断项目用户是否存在
// @Tags project-user
// @Param id path int true "项目ID"
// @Param user path int true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[bool]
// @Router /project/{id}/user/{user}/exists [get]
func projectUserExists(ctx *gin.Context) {
	pd := project.ProjectUser{
		ProjectId: ctx.Param("id"),
		UserId:    ctx.Param("user"),
	}
	has, err := db.Engine.Exist(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, has)
}

// @Summary 绑定项目用户
// @Schemes
// @Description 绑定项目用户
// @Tags project-user
// @Param id path int true "项目ID"
// @Param user path int true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /project/{id}/user/{user}/bind [get]
func projectUserBind(ctx *gin.Context) {
	pd := project.ProjectUser{
		ProjectId: ctx.Param("id"),
		UserId:    ctx.Param("user"),
	}
	_, err := db.Engine.InsertOne(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 删除项目用户
// @Schemes
// @Description 删除项目用户
// @Tags project-user
// @Param id path int true "项目ID"
// @Param user path int true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /project/{id}/user/{user}/unbind [get]
func projectUserUnbind(ctx *gin.Context) {
	_, err := db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("user")}).Delete(new(project.ProjectUser))
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 禁用项目用户
// @Schemes
// @Description 禁用项目用户
// @Tags project-user
// @Param id path int true "项目ID"
// @Param user path int true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /project/{id}/user/{user}/disable [get]
func projectUserDisable(ctx *gin.Context) {
	pd := project.ProjectUser{Disabled: true}
	_, err := db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("user")}).Cols("disabled").Update(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 启用项目用户
// @Schemes
// @Description 启用项目用户
// @Tags project-user
// @Param id path int true "项目ID"
// @Param user path int true "用户ID"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /project/{id}/user/{user}/enable [get]
func projectUserEnable(ctx *gin.Context) {
	pd := project.ProjectUser{Disabled: false}
	_, err := db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("user")}).Cols("disabled").Update(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

// @Summary 修改项目用户
// @Schemes
// @Description 修改项目用户
// @Tags project-user
// @Param id path int true "项目ID"
// @Param user path int true "用户ID"
// @Param project-user body project.ProjectUser true "项目用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} curd.ReplyData[int]
// @Router /project/{id}/user/{user} [post]
func projectUserUpdate(ctx *gin.Context) {
	var pd project.ProjectUser
	err := ctx.ShouldBindJSON(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	_, err = db.Engine.ID(schemas.PK{ctx.Param("id"), ctx.Param("user")}).
		Cols("user_id", "disabled").
		Update(&pd)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

func projectUserRouter(app *gin.RouterGroup) {
	app.GET("/list", projectUserList)
	app.GET("/:user/exists", projectUserExists)
	app.GET("/:user/bind", projectUserBind)
	app.GET("/:user/unbind", projectUserUnbind)
	app.GET("/:user/disable", projectUserDisable)
	app.GET("/:user/enable", projectUserEnable)
	app.POST("/:user", projectUserUpdate)
}
