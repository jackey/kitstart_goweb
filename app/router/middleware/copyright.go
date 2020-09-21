package middleware

import (
	"github.com/valyala/fasthttp"
)

// CopyrightMiddleware 版权中间件 - 显示版权信息
func CopyrightMiddleware(next MiddlewareFunc) MiddlewareFunc {
	return func(ctx *fasthttp.RequestCtx) {

		ctx.Response.Header.Set("server", "kitstart-goweb")

		next(ctx)

	}
}
