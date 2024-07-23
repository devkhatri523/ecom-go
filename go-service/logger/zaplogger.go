package logger

import (
	"context"
	"errors"
	"fmt"
	"github.com/devkhatri523/ecom-go/config/config"
	"github.com/devkhatri523/ecom-go/go-service/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

var zapSugarLogger *zap.SugaredLogger

func init() {
	logger, err := initLogger()
	if err != nil {
		log.Fatalf("Error while initializing zap logger. Error : %s", err)
	} else {
		zapSugarLogger = logger
	}

}

func initLogger() (*zap.SugaredLogger, error) {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	timeKey := defaultString(config.Default().GetString("logger.timeKey"), "timestamp")
	encoderConfig.TimeKey = timeKey
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC1123)
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "msg"
	encoderConfig.LevelKey = "level"
	encoderConfig.NameKey = "logger"
	encoderConfig.StacktraceKey = "trace"
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	defaultLogLevel := getDefaultLogLevel()
	consoleCore := getConsoleAppender(encoderConfig, defaultLogLevel)
	fileCore := getFileAppender(encoderConfig, defaultLogLevel)
	var cores []zapcore.Core
	if consoleCore != nil {
		cores = append(cores, consoleCore)
	}
	if fileCore != nil {
		cores = append(cores, fileCore)
	}
	if len(cores) > 0 {
		core := zapcore.NewTee(cores...)
		logger := zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel))
		return logger.Sugar(), nil
	} else {
		return nil, errors.New("could not create any logger")
	}
}
func getConsoleAppender(encoderConfig zapcore.EncoderConfig, defaultLogLevel zapcore.Level) zapcore.Core {
	isConsoleAppender := config.Default().GetBool("logger.consoleAppender")
	if isConsoleAppender {
		consoleAppender := zapcore.NewConsoleEncoder(encoderConfig)
		return zapcore.NewCore(consoleAppender, zapcore.AddSync(os.Stdout), defaultLogLevel)
	} else {
		return nil
	}
}

func getFileAppender(encoderConfig zapcore.EncoderConfig, defaultLogLevel zapcore.Level) zapcore.Core {
	isFileAppender := config.Default().GetBool("logger.fileAppender")
	if isFileAppender {
		fileAppender := zapcore.NewJSONEncoder(encoderConfig)
		logFileName := defaultString(config.Default().GetString("logger.fileName"), "log.json")
		logFile, _ := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		writer := zapcore.AddSync(logFile)
		return zapcore.NewCore(fileAppender, writer, defaultLogLevel)
	} else {
		return nil
	}
}

type ZapLogger struct {
}

func (z ZapLogger) Debug(args ...interface{}) {
	zapSugarLogger.Debug(args...)
}

func (z ZapLogger) Info(args ...interface{}) {
	zapSugarLogger.Info(args...)
}

func (z ZapLogger) Warn(args ...interface{}) {
	zapSugarLogger.Warn(args...)
}

func (z ZapLogger) Error(args ...interface{}) {
	zapSugarLogger.Error(args...)
}

func (z ZapLogger) Fatal(args ...interface{}) {
	zapSugarLogger.Fatal(args...)
}

func (z ZapLogger) DebugWithCtx(ctx context.Context, msg string) {
	zapSugarLogger.Debug(buildMsgWithCtx(ctx, msg))
}

func (z ZapLogger) InfoWithCtx(ctx context.Context, msg string) {
	zapSugarLogger.Info(buildMsgWithCtx(ctx, msg))
}

func (z ZapLogger) WarnWithCtx(ctx context.Context, msg string) {
	zapSugarLogger.Warn(buildMsgWithCtx(ctx, msg))
}

func (z ZapLogger) ErrorWithCtx(ctx context.Context, msg string) {
	zapSugarLogger.Error(buildMsgWithCtx(ctx, msg))
}

func (z ZapLogger) FatalWithCtx(ctx context.Context, msg string) {
	zapSugarLogger.Fatal(buildMsgWithCtx(ctx, msg))
}

func (z ZapLogger) Debugf(template string, args ...interface{}) {
	zapSugarLogger.Debugf(template, args...)
}

func (z ZapLogger) Infof(template string, args ...interface{}) {
	zapSugarLogger.Infof(template, args...)
}

func (z ZapLogger) Warnf(template string, args ...interface{}) {
	zapSugarLogger.Warnf(template, args...)
}

func (z ZapLogger) Errorf(template string, args ...interface{}) {
	zapSugarLogger.Errorf(template, args...)
}
func (z ZapLogger) Fatalf(template string, args ...interface{}) {
	zapSugarLogger.Fatalf(template, args...)
}

func (z ZapLogger) DebugWithCtxf(ctx context.Context, template string, args ...interface{}) {
	zapSugarLogger.Debugf(buildMsgWithCtx(ctx, template), args...)
}

func (z ZapLogger) InfoWithCtxf(ctx context.Context, template string, args ...interface{}) {
	zapSugarLogger.Infof(buildMsgWithCtx(ctx, template), args...)
}

func (z ZapLogger) WarnWithCtxf(ctx context.Context, template string, args ...interface{}) {
	zapSugarLogger.Warnf(buildMsgWithCtx(ctx, template), args...)
}

func (z ZapLogger) ErrorWithCtxf(ctx context.Context, template string, args ...interface{}) {
	zapSugarLogger.Errorf(buildMsgWithCtx(ctx, template), args...)
}

func (z ZapLogger) FatalWithCtxf(ctx context.Context, template string, args ...interface{}) {
	zapSugarLogger.Fatalf(buildMsgWithCtx(ctx, template), args...)
}

func buildMsgWithCtx(ctx context.Context, msg string) string {
	traceKey := trace.GetTraceKey()
	return fmt.Sprintf("%s %s : %s", msg, traceKey, ctx.Value(traceKey))
}
func getDefaultLogLevel() zapcore.Level {
	level := defaultString(config.Default().GetString("logger.logLevel"), "debug")
	switch level {
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
func defaultString(str string, defaultStr string) string {
	if str == "" {
		return defaultStr
	} else {
		return str
	}
}
