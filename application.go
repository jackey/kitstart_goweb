package main

import (
	"kitstart_goweb/app/utils"

	"kitstart_goweb/app/router"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
)

func main() {

	routerInstance := fasthttprouter.New()
	handler := router.InitRoutes(routerInstance)

	go fasthttp.ListenAndServe(":8085", handler)

	utils.Info("系统启动在 :8085")

	select {}

}
