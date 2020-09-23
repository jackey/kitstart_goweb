package router

import (
	"context"
	"kitstart_goweb/app/controller"
	"kitstart_goweb/app/utils"

	"github.com/valyala/fasthttprouter"
)

// WebRouter web 网站路由配置
type WebRouter struct {
	controllers []controller.BaseControllerInterface
}

// RouterInit web路由配置
func (webrouter *WebRouter) RouterInit(router *fasthttprouter.Router) {

	utils.Info("setup web routes")

	controller := controller.IndexController{
		controller.BaseController{
			Router: router,
		},
	}
	controller.SetupRoutes()
	webrouter.controllers = append(webrouter.controllers, &controller)

}

// InjectContext - 注入 context 到controllers
func (webrouter *WebRouter) InjectContext(ctx context.Context) {
	for _, contrl := range webrouter.controllers {
		contrl.SetContext(ctx)
	}
}
