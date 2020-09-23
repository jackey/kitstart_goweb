package middleware

import (
	"context"

	"github.com/valyala/fasthttp"
)

// CopyrightMiddleware 版权中间件 - 显示版权信息
func CopyrightMiddleware(referContext context.Context, next MiddlewareFunc) MiddlewareFunc {
	return func(referContext context.Context, ctx *fasthttp.RequestCtx) {

		ctx.Response.Header.Set("server", "kitstart-goweb")

		next(referContext, ctx)

	}
}
