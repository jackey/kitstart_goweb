package router

import (
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

	indexController := controller.NewIndexController()
	router.GET("/", indexController.Index)

}

// NewWebRouter 返回WebRouter
func NewWebRouter() *WebRouter {
	return &WebRouter{}
}
