package controller

import (
	"context"

	"github.com/valyala/fasthttprouter"
)

// BaseControllerInterface - 控制器接口
type BaseControllerInterface interface {
	SetupRoutes()                   // 配置 routes
	SetContext(ctx context.Context) // 注入 context
}

// BaseController 控制器基类
type BaseController struct {
	Router  *fasthttprouter.Router
	Context context.Context
}

// SetContext - 注入 context
func (controller *BaseController) SetContext(ctx context.Context) {
	controller.Context = ctx
}
