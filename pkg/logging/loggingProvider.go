package logging

import (
	"os"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerKeyType string

const correlationIdKey loggerKeyType = "correlationId"

type LogProvider struct {
	ZapLogger     *zap.Logger
	CorrelationId uuid.UUID
}

var logger *LogProvider

func InitLogger(debug bool) *LogProvider {
	config := zap.NewProductionConfig()
	if debug {
		config = zap.NewDevelopmentConfig()
	}
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339) // or time.RubyDate or "2006-01-02 15:04:05" or even freaking time.Kitchen
	zapLogger, err := config.Build()

	if err != nil {
		panic("Logger could not initialized")
	}

	// zap.ReplaceGlobals(zapLogger)
	logger = &LogProvider{
		ZapLogger: zapLogger,
	}
	return logger
}

func GetLogger() *LogProvider {
	if logger == nil {
		InitLogger(os.Getenv("ENVIRONMENT") == "DEV")
	}
	return logger
}

func (l *LogProvider) Sync() {
	l.ZapLogger.Sync()
}

func (l *LogProvider) Exception(message string, err error) {
	l.ZapLogger.With(
		// zap.Time("DateTime", time.Now()),
		// zap.NamedError("Exception", err),
		zap.String(string(correlationIdKey), l.CorrelationId.String())).Error(message)
}

func (l *LogProvider) GenerateCorrelationId() uuid.UUID {
	l.CorrelationId = uuid.New()
	return l.CorrelationId
}

func (l *LogProvider) Info(message string) {
	l.ZapLogger.With(
		zap.String(string(correlationIdKey), l.CorrelationId.String())).Info(message)
}

func (l *LogProvider) Error(message string, err error) {
	l.ZapLogger.With(
		zap.String(string(correlationIdKey), l.CorrelationId.String())).Error(message, zap.Error(err))
}
