package router

import (
	"kitstart_goweb/app/router/middleware"
	"kitstart_goweb/app/utils"

	"github.com/valyala/fasthttp"

	"github.com/valyala/fasthttprouter"
)

// Setuper 路由初始化接口
type Setuper interface {
	RouterInit(router *fasthttprouter.Router)
}

var routerMap = make(map[string]Setuper)

// 模块初始化 - 注册路由
func init() {
	routerMap["webRouter"] = NewWebRouter()
	routerMap["apiRouter"] = &APIRouter{}
}

// InitRoutes  初始化所有的路由配置
func InitRoutes(router *fasthttprouter.Router) func(ctx *fasthttp.RequestCtx) {

	for k, routes := range routerMap {
		utils.Info("注册 %s 路由 ", k)
		routes.RouterInit(router)
	}

	// 配置中间件
	handler := middleware.CopyrightMiddleware(router.Handler)

	return handler

}
