package controller

import (
	"context"
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

// IndexController 网站主入口 控制器
type IndexController struct {
	context context.Context
}

// Index - 首页
func (controller *IndexController) Index(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	fmt.Fprintf(ctx, "i'm index action ")
}

// NewIndexController - 新建控制器实例
func NewIndexController() *IndexController {
	return &IndexController{
		context: context.Background(),
	}
}
