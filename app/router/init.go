package router

import (
	"context"
	"kitstart_goweb/app/router/middleware"

	"github.com/valyala/fasthttp"

	"github.com/valyala/fasthttprouter"
)

// Setuper 路由初始化接口
type Setuper interface {
	RouterInit(router *fasthttprouter.Router)
	InjectContext(ctx context.Context)
}

var routerMap = make(map[string]Setuper)

// 模块初始化 - 注册路由
func init() {
	routerMap["webRouter"] = &WebRouter{}
	routerMap["apiRouter"] = &APIRouter{}
}

// InitRoutes  初始化所有的路由配置
func InitRoutes(router *fasthttprouter.Router) func(ctx *fasthttp.RequestCtx) {

	referContext := context.Background()

	for _, routes := range routerMap {
		routes.RouterInit(router)
	}

	defaultHandler := func(referCon context.Context, ctx *fasthttp.RequestCtx) {

		// 注入 context 到控制器
		for _, routes := range routerMap {
			routes.InjectContext(referCon)
		}

		router.Handler(ctx)
	}

	// 配置中间件
	handler := middleware.CopyrightMiddleware(referContext, defaultHandler)
	handler = middleware.Logger(referContext, handler)

	// 传入 context 参数
	return func(ctx *fasthttp.RequestCtx) {
		handler(referContext, ctx)
	}

}
