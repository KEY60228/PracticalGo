package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
)

type contextKey string

const logKey contextKey = "log"

func GetLogger(ctx context.Context) zerolog.Logger {
	return ctx.Value(logKey).(zerolog.Logger)
}

type responseWriterWrapper struct {
	status       int
	responseSize int
	writer       http.ResponseWriter
	request      *http.Request
	start        time.Time
}

func (r *responseWriterWrapper) Flush() {
	flusher := r.writer.(http.Flusher)
	flusher.Flush()
	r.status = 200
}

func (r responseWriterWrapper) Header() http.Header {
	return r.writer.Header()
}

func (r responseWriterWrapper) Write(content []byte) (int, error) {
	r.responseSize += len(content)
	if r.status == 0 {
		r.status = http.StatusOK
	}
	return r.writer.Write(content)
}

func (r *responseWriterWrapper) WriteHeader(statusCode int) {
	r.status = statusCode
	r.writer.WriteHeader(statusCode)
}

func (r *responseWriterWrapper) MarshalZerologObject(e *zerolog.Event) {
	e.Str("requestMethod", r.request.Method)
	e.Str("requestURL", r.request.URL.String())
	e.Int64("requestSize", r.request.ContentLength)
	e.Int("status", r.status)
	e.Int("responseSize", r.responseSize)
	e.Str("referer", r.request.Header.Get("Referer"))
	e.Str("latency", time.Now().Sub(r.start).String())
	e.Bool("cacheHit", r.status == 304)
	forwarded := r.request.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		e.Str("remoteIP", forwarded)
	} else {
		e.Str("protocol", r.request.Proto)
	}
}

var _ http.ResponseWriter = &responseWriterWrapper{}
var _ http.Flusher = &responseWriterWrapper{}
var _ zerolog.LogObjectMarshaler = &responseWriterWrapper{}

func WithLogger(logger zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			traceID := r.Header.Get("Trace-ID")
			if traceID == "" {
				traceID = xid.New().String()
			}
			logger = logger.With().Str("trace-id", traceID).Logger()
			ctx := context.WithValue(r.Context(), logKey, logger)

			writer := &responseWriterWrapper{
				writer:  w,
				request: r,
				start:   time.Now(),
			}
			next.ServeHTTP(writer, r.WithContext(ctx))
			logger.Info().Object("httpRequest", writer).Send()
		})
	}
}

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	r := chi.NewRouter()
	r.Use(WithLogger(logger))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World")
	})

	log.Fatal(http.ListenAndServe("localhost:8888", r))
}
