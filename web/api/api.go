package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/storm/v3"
	"github.com/zgwit/storm/v3/q"
	"net/http"
	"reflect"
)

type paramSearchEx struct {
	Skip     int                      `form:"skip"`
	Limit    int                      `form:"limit"`
	Sort     map[string]int           `form:"sort"`
	Filters  map[string][]interface{} `form:"filter"`
	Keywords map[string]string        `json:"keyword"`
}

type paramID struct {
	ID int `uri:"id"`
}

type WatchMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
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

func parseParamId(ctx *gin.Context) {
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
	app.GET("/logout", logout)

	//检查 session，必须登录
	app.Use(mustLogin)

	projectRoutes(app.Group("/project"))
	deviceRoutes(app.Group("/device"))
	elementRoutes(app.Group("/element"))
	templateRoutes(app.Group("/template"))
	tunnelRoutes(app.Group("/tunnel"))
	linkRoutes(app.Group("/link"))
	userRoutes(app.Group("/user"))

	//TODO 报接口错误（以下代码不生效，路由好像不是树形处理）
	app.Use(func(ctx *gin.Context) {
		replyFail(ctx, "Not found")
		ctx.Abort()
	})
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

func normalSearch(ctx *gin.Context, store storm.Node, to interface{}) (interface{}, int, error) {
	var body paramSearchEx
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		return nil, 0, err
	}

	cond := make([]q.Matcher, 0)

	//过滤
	for k, v := range body.Filters {
		if len(v) > 0 {
			if len(v) == 1 {
				cond = append(cond, q.Eq(k, v[0]))
			} else {
				cond = append(cond, q.In(k, v))
			}
		}
	}

	//关键字搜索
	kws := make([]q.Matcher, 0)
	for k, v := range body.Keywords {
		if v != "" {
			kws = append(kws, q.Re(k, v))
		}
	}
	if len(kws) > 0 {
		cond = append(cond, q.Or(kws...))
	}

	query := store.Select(cond...)

	//查询总数
	cnt, err := query.Count(to)
	if err != nil {
		return nil, 0, err
	}
	//cnt := 0

	//分页
	query = query.Skip(body.Skip).Limit(body.Limit)

	//排序
	if len(body.Sort) > 0 {
		for k, v := range body.Sort {
			if v > 0 {
				query = query.OrderBy(k)
			} else {
				query = query.OrderBy(k).Reverse()
			}
		}
	} else {
		query = query.OrderBy("ID").Reverse()
	}

	//查询
	//res := reflect.MakeSlice(reflect.TypeOf(to), 0, body.Length).Interface()
	res := reflect.New(reflect.SliceOf(reflect.TypeOf(to))).Interface()
	err = query.Find(res)
	if err != nil && err != storm.ErrNotFound {
		return nil, cnt, err
	}

	return res, cnt, nil
}
