package main

import (
	_ "cleverbamboo.com/bee-shop-b2c/routers"
	_ "cleverbamboo.com/bee-shop-b2c/sysinit"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
)

var success = []byte("SUPPORT OPTIONS")

var corsFunc = func(ctx *context.Context) {
	origin := ctx.Input.Header("Origin")
	ctx.Output.Header("Access-Control-Allow-Methods", "OPTIONS,DELETE,POST,GET,PUT,PATCH")
	ctx.Output.Header("Access-Control-Max-Age", "3600")
	ctx.Output.Header("Access-Control-Allow-Headers", "X-Custom-Header,accept,Content-Type,Access-Token")
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	ctx.Output.Header("Access-Control-Allow-Origin", origin)
	if ctx.Input.Method() == http.MethodOptions {
		// options请求，返回200
		ctx.Output.SetStatus(http.StatusOK)
		_ = ctx.Output.Body(success)
	}
}


func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, corsFunc)

	beego.Run()
}
