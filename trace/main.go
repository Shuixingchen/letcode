package main

import (
	"io"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// NewTracer 创建一个jaeger Tracer
func CreateTracer(serviceName, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{ // 采样器配置
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{ // 报告器配置
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			CollectorEndpoint:   addr,
		},
	}
	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))

	return tracer, closer, err
}

func main() {
	tracer, closer, _ := CreateTracer("UserinfoService", "http://127.0.0.1:14268/api/traces")
	// 创建第一个 span A
	parentSpan := tracer.StartSpan("A")
	// 调用其它服务
	GetUserInfo(tracer, parentSpan)
	// 结束 A
	parentSpan.Finish()
	// 结束当前 tracer
	closer.Close()
}

// 请求远程服务，获得用户信息
func GetUserInfo(tracer opentracing.Tracer, parentSpan opentracing.Span) {
	// 继承上下文关系，创建子 span
	childSpan := tracer.StartSpan(
		"B",
		opentracing.ChildOf(parentSpan.Context()),
	)

	url := "http://127.0.0.1:8888"
	req, _ := http.NewRequest("GET", url, nil)
	// 设置 tag，这个 tag 我们后面讲
	ext.SpanKindRPCClient.Set(childSpan)
	ext.HTTPUrl.Set(childSpan, url)
	ext.HTTPMethod.Set(childSpan, "GET")
	tracer.Inject(childSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	resp, _ := http.DefaultClient.Do(req)
	_ = resp // 丢掉
	defer childSpan.Finish()
}
