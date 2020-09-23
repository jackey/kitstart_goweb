package middleware

import (
	"context"

	"kitstart_goweb/app/utils"

	"github.com/valyala/fasthttp"
)

// Logger - 记录http的访问日志
func Logger(referContext context.Context, next MiddlewareFunc) MiddlewareFunc {
	return func(referContext context.Context, ctx *fasthttp.RequestCtx) {

		go utils.Info(ctx.RemoteAddr(), "-", string(ctx.Method()), "-", string(ctx.Path()))

		next(referContext, ctx)
	}
}
