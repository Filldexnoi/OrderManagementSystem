package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

var Tracer = otel.GetTracerProvider().Tracer("fiber-server")

func TracingMiddleware(c *fiber.Ctx) error {
	carrier := propagation.HeaderCarrier(c.GetReqHeaders())
	ctx := otel.GetTextMapPropagator().Extract(c.Context(), carrier)

	spanOptions := []trace.SpanStartOption{
		trace.WithAttributes(semconv.HTTPMethodKey.String(c.Method())),
		trace.WithAttributes(semconv.HTTPTargetKey.String(c.Path())),
		trace.WithAttributes(semconv.HTTPRouteKey.String(c.Path())),
		trace.WithAttributes(semconv.HTTPURLKey.String(fmt.Sprintf("%s://%s%s", c.Protocol(), c.Hostname(), c.OriginalURL()))),
		trace.WithAttributes(semconv.UserAgentOriginalKey.String(c.Get("User-Agent"))),
		trace.WithAttributes(semconv.HTTPRequestContentLengthKey.Int(c.Request().Header.ContentLength())),
		trace.WithAttributes(semconv.HTTPSchemeKey.String(c.Protocol())),
		trace.WithAttributes(semconv.NetTransportTCP),
		trace.WithSpanKind(trace.SpanKindServer),
	}

	ctx, span := Tracer.Start(ctx, c.Method()+" "+c.Path(), spanOptions...)
	defer span.End()

	c.SetUserContext(ctx)

	if err := c.Next(); err != nil {
		c.Response().SetStatusCode(fiber.StatusInternalServerError)
		return err
	}

	propagator := otel.GetTextMapPropagator()
	carrier = propagation.HeaderCarrier{}
	propagator.Inject(ctx, carrier)

	for _, k := range carrier.Keys() {
		c.Response().Header.Set(k, carrier.Get(k))
	}

	span.SetAttributes(semconv.HTTPStatusCodeKey.Int(c.Response().StatusCode()))
	return nil
}
