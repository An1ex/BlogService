package tracer

import (
	"io"
	"time"

	"BlogService/global"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, nil
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}

func InitTracer(serviceName, agentHostPort string) error {
	jaegerTracer, _, err := NewJaegerTracer(serviceName, agentHostPort)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
