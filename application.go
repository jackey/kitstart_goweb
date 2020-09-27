package main

import (
	"kitstart_goweb/app"
	"kitstart_goweb/app/utils"
)

func main() {

	application := app.InitApp()

	utils.Info("app listening at :", application.Port)

	select {}
}
