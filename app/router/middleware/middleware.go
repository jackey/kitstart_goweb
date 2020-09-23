package middleware

import (
	"context"

	"github.com/valyala/fasthttp"
)

// MiddlewareFunc 中间件类型
type MiddlewareFunc func(context context.Context, ctx *fasthttp.RequestCtx)
