package router

import (
	"kitstart_goweb/app/utils"

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
}

// InitRoutes  初始化所有的路由配置
func InitRoutes(router *fasthttprouter.Router) {
	for k, routes := range routerMap {
		utils.Info("注册 %s 路由 ", k)
		routes.RouterInit(router)
	}
}
