package middleware

import (
	"github.com/valyala/fasthttp"
)

// MiddlewareFunc 中间件类型
type MiddlewareFunc func(ctx *fasthttp.RequestCtx)
