package middleware

import (
	"bytes"
	"dca-bot-live/app/utils"
	"io"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func JaegerTrace() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			ct := req.Context()

			reqBody := utils.LogRequestWriter(req)
			req.Body = io.NopCloser(bytes.NewReader(reqBody))

			// === Start tracing ===
			tracer := otel.Tracer("dca-bot-live")
			ctx, span := tracer.Start(ct, req.Method+" "+req.URL.Path)
			defer span.End()
			c.SetRequest(req.WithContext(ctx)) // with tracing context

			// === Wrap response ===
			writer, respBuffer := utils.LogResponseWriter(res.Writer)
			res.Writer = writer

			start := time.Now()
			err := next(c) // handler
			stop := time.Now()

			// === Add span attributes ===
			span.AddEvent("Incoming HTTP request", trace.WithAttributes(
				attribute.String("method", req.Method),
				attribute.String("path", req.URL.Path),
				attribute.Int("status", res.Status),
				attribute.String("latency", stop.Sub(start).String()),
				attribute.String("ip", c.RealIP()),
				attribute.String("request_body", strings.TrimSpace(string(reqBody))),
				attribute.String("response_body", strings.TrimSpace(respBuffer.String())),
			))

			return err
		}
	}
}
