package middleware

import (
	"letcode/mygin/trace"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// 其他进程通过 HTTP 调用 当前webserver时，通过 Http Header 携带 trace 信息(称为上下文)，这个中间件就是解析trace信息,并创建新tracer.

func UseOpenTracing() gin.HandlerFunc {
	handler := func(c *gin.Context) {
		tracer, spanContext, closer, _ := trace.CreateTracer("WebService", "http://127.0.0.1:14268/api/traces", c.Request.Header)
		defer closer.Close()
		// 生成依赖关系，并新建一个 span、
		// 这里很重要，因为生成了  References []SpanReference 依赖关系
		startSpan := tracer.StartSpan(c.Request.URL.Path, ext.RPCServerOption(spanContext))
		defer startSpan.Finish()

		// 记录 tag, 用于日志查询过滤
		ext.HTTPUrl.Set(startSpan, c.Request.URL.Path)
		ext.HTTPMethod.Set(startSpan, c.Request.Method)
		ext.Component.Set(startSpan, "Gin-Http")

		// 在 header 中加上当前进程的上下文信息
		c.Request = c.Request.WithContext(opentracing.ContextWithSpan(c.Request.Context(), startSpan))
		// 传递给下一个中间件
		c.Next()
		// 继续设置 tag
		ext.HTTPStatusCode.Set(startSpan, uint16(c.Writer.Status()))
	}

	return handler
}
