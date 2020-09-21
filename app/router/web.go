package router

import (
	"kitstart_goweb/app/utils"

	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

// WebRouter web 网站路由配置
type WebRouter struct {
}

// RouterInit web路由配置
func (webrouter *WebRouter) RouterInit(router *fasthttprouter.Router) {
	utils.Info("注册web路由列表")

	router.GET("/", func(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
		fmt.Fprintln(ctx, "Welcome \n")
	})

}

// NewWebRouter 返回WebRouter
func NewWebRouter() *WebRouter {
	return &WebRouter{}
}
