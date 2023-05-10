package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	date := time.Now().Format("2006-01-02")
	appName := "service"
	filename := fmt.Sprintf("logs/%v.log", appName+"."+date)
	//filename := "logs.log"
	Logger = FileLogger(filename)
}

func FileLogger(filename string) *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the HTTP request details
		start := time.Now()
		reqId := c.Writer.Header().Get("X-Request-Id")
		if reqId == "" {
			reqId = strconv.FormatInt(time.Now().UnixNano(), 36)
			c.Writer.Header().Set("X-Request-Id", reqId)
		}

		headers := getRequestHeaders(c.Request)
		reqBody := getRequestBytes(c.Request)

		// Create a response writer proxy to capture the response body
		writer := &responseWriter{c.Writer, http.StatusOK, bytes.Buffer{}}
		// Store the response writer proxy in the gin context
		c.Writer = writer

		statusCode := c.Writer.Status()
		path := c.Request.URL.Path
		method := c.Request.Method
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		latency := time.Since(start)
		referer := c.Request.Referer()

		// Create a logger using the Zap library
		requestLogger := logger.With(
			zap.String("id", reqId),
			zap.Int("status_code", statusCode),
			zap.String("path", path),
			zap.String("method", method),
			zap.String("ip", clientIP),
			zap.String("user_agent", userAgent),
			zap.String("referer", referer),
			zap.Duration("latency", latency),
		)

		// Log the request using the logger
		requestLogger.Info("Request",
			zap.Any("Headers", headers),
			zap.ByteString("requestBody", reqBody))

		defer func() {
			if err := recover(); err != nil {
				requestLogger.Error("Panic occurred",
					zap.Any("error", err),
					zap.String("stacktrace", string(debug.Stack())))
			}
		}()

		// Call the next middleware or handler
		c.Next()

		// Log the response using the logger
		responseBody := writer.body.String()
		requestLogger.Info("Response",
			zap.ByteString("responseBody", []byte(responseBody)))

	}
}

func getRequestHeaders(r *http.Request) map[string]string {
	headers := make(map[string]string)
	for name, values := range r.Header {
		for _, value := range values {
			headers[name] = value
		}
	}
	return headers
}

// Helper function to get the request body as a byte string
func getRequestBytes(r *http.Request) []byte {
	if r.Body == nil {
		return []byte{}
	}
	defer r.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes
}

// Custom response writer proxy to capture the response body
type responseWriter struct {
	gin.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (w *responseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	n, err := w.body.Write(b)
	if err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(b)
}
