package mtl

import (
	"common/consts"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/kitex/server"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

var TracerProvider *tracesdk.TracerProvider

/**
* Tracing 提供了从请求开始接收到处理完毕的整个生命周期的全貌
* https://cloudwego.cn/zh/docs/hertz/tutorials/third-party/open-telemetry/#tracing
 */

func InitTracing(serviceName string) {

	ctx := context.Background()
	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint(consts.OpenTelemetryAddress),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String(serviceName)),
	)
	if err != nil {
		panic(fmt.Errorf("build resource: %w", err))
	}
	processor := tracesdk.NewBatchSpanProcessor(exporter)

	TracerProvider = tracesdk.NewTracerProvider(
		tracesdk.WithSpanProcessor(processor),
		tracesdk.WithResource(res),
	)

	otel.SetTracerProvider(TracerProvider)

	server.RegisterShutdownHook(func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := exporter.Shutdown(ctx); err != nil {
			log.Printf("exporter shutdown: %v", err)
		}
	})
}
