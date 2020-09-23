package router

import (
	"context"
	"kitstart_goweb/app/utils"

	"kitstart_goweb/app/controller"

	"github.com/valyala/fasthttprouter"
)

// APIRouter API路由配置
type APIRouter struct {
	controllers []controller.BaseControllerInterface
}

// RouterInit 初始化路由
func (api *APIRouter) RouterInit(router *fasthttprouter.Router) {
	utils.Info("init api routes")

}

// InjectContext - 注入 context 到 controller
func (api *APIRouter) InjectContext(ctx context.Context) {
	for _, controller := range api.controllers {
		controller.SetContext(ctx)
	}
}
