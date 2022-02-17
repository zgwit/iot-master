package api

import (
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type paramFilter struct {
	Key    string        `form:"key"`
	Values []interface{} `form:"value"`
}

type paramKeyword struct {
	Key   string `form:"key"`
	Value string `json:"value"`
}

type paramSearch struct {
	Offset    int            `form:"offset"`
	Length    int            `form:"length"`
	SortKey   string         `form:"sortKey"`
	SortOrder string         `form:"sortOrder"`
	Filters   []paramFilter  `form:"filters"`
	Keywords  []paramKeyword `json:"keywords"`
	//Keyword   string        `form:"keyword"`
}

type paramID struct {
	ID int `uri:"id"`
}

func mustLogin(ctx *gin.Context) {
	session := sessions.Default(ctx)
	if user := session.Get("user"); user != nil {
		ctx.Set("user", user)
		ctx.Next()
	} else {
		//TODO 检查OAuth2返回的code，进一步获取用户信息，放置到session中

		ctx.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "Unauthorized"})
		ctx.Abort()
	}
}

func parseParamId(ctx *gin.Context)  {
	var pid paramID
	err := ctx.ShouldBindUri(&pid)
	if err != nil {
		replyError(ctx, err)
		ctx.Abort()
		return
	}
	ctx.Set("id", pid.ID)
	ctx.Next()
}

func RegisterRoutes(app *gin.RouterGroup) {
	app.POST("/login", login)
	app.Any("/logout", logout)

	//检查 session，必须登录
	app.Use(mustLogin)

	projectRoutes(app.Group("/project"))

	//TODO 转移至子目录，并使用中间件，检查session及权限
	//mod := reflect.TypeOf(model.Tunnel{})
	////app.POST("/project/:id/tunnels", curdApiListByID(mod, "project_id"))
	//app.POST("/tunnels", curdApiList(mod))
	//app.POST("/tunnel", curdApiCreate(mod, nil, nil))                  //TODO 启动
	//app.DELETE("/tunnel/:id", curdApiDelete(mod, nil, nil))            //TODO 停止
	//app.PUT("/tunnel/:id", curdApiModify(mod, []string{""}, nil, nil)) //TODO 重新启动
	//app.GET("/tunnel/:id", curdApiGet(mod))
	//
	//app.GET("/tunnel/:id/start", tunnelStart)
	//app.GET("/tunnel/:id/stop", tunnelStop)

	//app.POST("/channel/:id/links")

	//连接管理
	//mod = reflect.TypeOf(model.Link{})
	////app.POST("/tunnel/:id/links", curdApiListByID(mod, "tunnel_id"))
	//app.POST("/links", curdApiList(mod))
	//app.DELETE("/link/:id", curdApiDelete(mod, nil, nil)) //TODO 停止
	//app.PUT("/link/:id", curdApiModify(mod, []string{""}, nil, nil))
	//app.GET("/link/:id", curdApiGet(mod))
	//
	////设备管理
	//mod = reflect.TypeOf(model.Device{})
	////app.POST("/project/:id/devices", curdApiListByID(mod, "project_id"))
	//app.POST("/devices", curdApiList(mod))
	//app.POST("/device", curdApiCreate(mod, nil, nil))
	//app.DELETE("/device/:id", curdApiDelete(mod, nil, nil))
	//app.PUT("/device/:id", curdApiModify(mod, []string{""}, nil, nil))
	//app.GET("/device/:id", curdApiGet(mod))
	//
	////项目管理
	//mod = reflect.TypeOf(model.Project{})
	//app.POST("/projects", curdApiList(mod))
	//app.POST("/project", curdApiCreate(mod, projectBeforeCreate, projectAfterCreate))
	//app.DELETE("/project/:id", curdApiDelete(mod, nil, projectAfterDelete))
	//app.PUT("/project/:id", curdApiModify(mod, []string{""}, nil, projectAfterModify))
	//app.GET("/project/:id", curdApiGet(mod))

}

func replyList(ctx *gin.Context, data interface{}, total int) {
	ctx.JSON(http.StatusOK, gin.H{"data": data, "total": total})
}

func replyOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func replyFail(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusOK, gin.H{"error": err})
}

func replyError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
}

func nop(ctx *gin.Context) {
	ctx.String(http.StatusForbidden, "Unsupported")
}

func normalSearch(ctx *gin.Context, store storm.Node, to interface{}) (int, error) {
	var body paramSearch
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		return 0, err
	}

	cond := make([]q.Matcher, 0)

	//过滤
	for _, filter := range body.Filters {
		if len(filter.Values) > 0 {
			if len(filter.Values) == 1 {
				cond = append(cond, q.Eq(filter.Key, filter.Values[0]))
			} else {
				cond = append(cond, q.In(filter.Key, filter.Values))
			}
		}
	}

	//关键字搜索
	kws := make([]q.Matcher, 0)
	for _, keyword := range body.Keywords {
		if keyword.Value != "" {
			kws = append(kws, q.Re(keyword.Key, keyword.Value))
		}
	}
	if len(kws) > 0 {
		cond = append(cond, q.Or(kws...))
	}

	query := store.Select(cond...)

	//查询总数
	cnt, err := query.Count(to)
	if err != nil {
		return cnt, err
	}

	//分页
	query = query.Skip(body.Offset).Limit(body.Length)

	//排序
	if body.SortKey != "" {
		if body.SortOrder == "desc" {
			query = query.OrderBy(body.SortKey).Reverse()
		} else {
			query = query.OrderBy(body.SortKey)
		}
	} else {
		query = query.OrderBy("ID").Reverse()
	}

	//查询
	err = query.Find(to)
	if err != nil && err != storm.ErrNotFound {
		return 0, err
	}

	return cnt, nil
}
