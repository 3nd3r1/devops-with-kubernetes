package middleware

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"todo-project-backend/internal/logger"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buffer, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(buffer))
		requestBodyString := string(buffer)

		logger.Logger.Info("incoming request",
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.RequestURI),
			slog.String("body", requestBodyString),
			slog.String("referrer", c.Request.Referer()),
		)

		c.Next()

	}
}

type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w ResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		responseBody := new(bytes.Buffer)
		writer := ResponseWriter{body: responseBody, ResponseWriter: c.Writer}
		c.Writer = writer

		c.Next()

		logger.Logger.Info("outgoing response",
			slog.Int("status", c.Writer.Status()),
			slog.String("path", c.Request.RequestURI),
			slog.String("request_id", c.Writer.Header().Get("Request-Id")),
			slog.String("body", responseBody.String()),
		)

		if c.Writer.Status() == http.StatusInternalServerError {
			logger.Logger.Error(c.Errors.String())
		}
	}
}
