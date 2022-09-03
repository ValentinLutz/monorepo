package api

import (
	"app/api/order"
	"app/internal/errors"
	"app/internal/util"
	"bytes"
	"io"
	"net/http"
	"time"
)

// RequestResponseLogger is a middleware handler that does log requests and responses
// when a client or server error occurs
type RequestResponseLogger struct {
	handler http.Handler
	logger  *util.Logger
}

func (rrl *RequestResponseLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		rrl.logger.Log().Error().Err(err).Msg("Error reading request body")
		orderapi.Error(w, r, http.StatusInternalServerError, errors.Panic, err.Error())
		return
	}
	reader := io.NopCloser(bytes.NewBuffer(requestBody))
	r.Body = reader

	rw := wrapResponseWriter(w)

	rrl.handler.ServeHTTP(rw, r)

	if rw.status >= 400 {
		rrl.logRequest(r, requestBody)
		rrl.logResponse(startTime, rw)
	}
}

func (rrl *RequestResponseLogger) logRequest(r *http.Request, requestBody []byte) {
	rrl.logger.Log().Info().
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Str("body", string(requestBody)).
		Msg("Request")
}

func (rrl *RequestResponseLogger) logResponse(startTime time.Time, rw *responseWriter) {
	rrl.logger.Log().Info().
		Str("duration", time.Since(startTime).String()).
		Int("status", rw.status).
		Str("body", string(rw.body)).
		Msg("Response")
}

func NewRequestResponseLogger(handlerToWrap http.Handler, logger *util.Logger) *RequestResponseLogger {
	return &RequestResponseLogger{
		handler: handlerToWrap,
		logger:  logger,
	}
}

// responseWriter is a wrapper for http.ResponseWriter that allows
// the written HTTP status code and
// the written HTTP body to be captured for logging.
type responseWriter struct {
	http.ResponseWriter
	status int
	body   []byte
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body = b
	return rw.ResponseWriter.Write(b)
}
