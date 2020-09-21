package router

import (
	"fmt"
	"kitstart_goweb/app/utils"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

// APIRouter API路由配置
type APIRouter struct{}

// RouterInit 初始化路由
func (api *APIRouter) RouterInit(router *fasthttprouter.Router) {

	utils.Info("初始化 API 路由信息")

	router.GET("/api", func(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
		fmt.Fprintf(ctx, "api version %s \n", "0.0.1.beta")
	})

}
