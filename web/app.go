package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/zgwit/iot-master/v3/internal"
	"io"
	"net"
	"net/http"
	"net/url"
)

func appProxy(ctx *gin.Context) {
	svc := internal.Applications.Load(ctx.Param("id"))
	if svc == nil {
		_ = ctx.Error(errors.New("应用未注册"))
		return
	}

	//省去前缀
	//l := len("/service/" + ctx.Param("name"))
	//u := ctx.Request.RequestURI[l:]

	req := ctx.Request.Clone(ctx)
	req.URL, _ = url.Parse(svc.Address + req.RequestURI)
	req.RequestURI = ""

	//req, _ := http.NewRequest(ctx.Request.Method, ctx.Request.RequestURI, ctx.Request.Body)
	cli := http.Client{}
	if svc.Type == "unix" {
		cli.Transport = &http.Transport{DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.Dial(network, addr)
		}}
	}

	resp, err := cli.Do(req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	//返回结果
	//ctx.Request.Response = resp
	//_ = resp.Write(ctx.Writer)
	_, err = io.Copy(ctx.Writer, resp.Body)
	if err != nil {
		_ = ctx.Error(errors.New("应用未注册"))
		return
	}

	ctx.Abort()
}
