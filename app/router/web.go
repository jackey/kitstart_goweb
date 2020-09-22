package router

import (
	"context"
	"kitstart_goweb/app/controller"
	"kitstart_goweb/app/utils"

	"github.com/valyala/fasthttprouter"
)

// WebRouter web 网站路由配置
type WebRouter struct {
}

// RouterInit web路由配置
func (webrouter *WebRouter) RouterInit(router *fasthttprouter.Router) {

	utils.Info("注册web路由列表")

	controller := controller.IndexController{
		controller.BaseController{
			Router:  router,
			Context: context.Background(),
		},
	}
	controller.SetupRoutes()

}

// NewWebRouter 返回WebRouter
func NewWebRouter() *WebRouter {
	return &WebRouter{}
}
