package controller

import (
	"context"

	"github.com/valyala/fasthttprouter"
)

type controllerInterface interface {
	SetupRoutes()
}

// BaseController 控制器基类
type BaseController struct {
	Router  *fasthttprouter.Router
	Context context.Context
}
