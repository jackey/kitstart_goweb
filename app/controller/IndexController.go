package controller

import (
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

// IndexController 网站主入口 控制器
type IndexController struct {
	BaseController
}

// Index - 首页
func (controller *IndexController) Index(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	fmt.Fprintf(ctx, "i'm index action and data is \n")
}

// SetupRoutes - 配置路由表
func (controller *IndexController) SetupRoutes() {
	controller.Router.GET("/", controller.Index)
}
