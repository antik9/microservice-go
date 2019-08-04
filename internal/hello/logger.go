package hello

import (
	"log"
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func getLevel() zapcore.Level {
	switch level := ServerConfig.Log.Level; level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func init() {
	var err error
	logger, err = zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(getLevel()),
		OutputPaths: []string{ServerConfig.Log.File},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "server",
		},
	}.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
}

func LogResponse(r *http.Request, status int) {
	logger.Info(ServerConfig.Web.Name,
		zap.String("proto", r.Proto),
		zap.String("method", r.Method),
		zap.String("from", r.RemoteAddr),
		zap.String("url", r.URL.String()),
		zap.Int("status", status))
}
