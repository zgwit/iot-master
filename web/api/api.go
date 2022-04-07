package api

import (
	"github.com/zgwit/storm/v3"
	"github.com/zgwit/storm/v3/q"
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
