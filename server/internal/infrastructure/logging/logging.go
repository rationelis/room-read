package logging

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"room_read/internal/infrastructure/configuration"
)

const (
	TraceIdKey = "traceId"
)

func SetupLogger(configuration configuration.Configuration) error {
	level, err := parseConfigurationLogLevel(configuration.Logging.LogLevel)

	if err != nil {
		return err
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	slog.SetDefault(logger)

	return nil
}

func parseConfigurationLogLevel(level string) (slog.Level, error) {
	switch level {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warning":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return 0, errors.New(fmt.Sprintf("Invalid log level [%s], valid levels are [error, warning, info, debug]", level))
	}
}

func Info(ctx context.Context, message string, args ...slog.Attr) {
	doLog(ctx, slog.LevelInfo, message, args...)
}

func Debug(ctx context.Context, message string, args ...slog.Attr) {
	doLog(ctx, slog.LevelDebug, message, args...)
}

func Error(ctx context.Context, message string, args ...slog.Attr) {
	doLog(ctx, slog.LevelError, message, args...)
}

func WithError(ctx context.Context, err error, args ...slog.Attr) {
	doLog(ctx, slog.LevelError, err.Error(), append(args)...)
}

func Warn(ctx context.Context, message string, args ...slog.Attr) {
	doLog(ctx, slog.LevelWarn, message, args...)
}

func doLog(ctx context.Context, logLevel slog.Level, message string, args ...slog.Attr) {
	if ctx.Value(TraceIdKey) == nil {
		ctx = context.WithValue(ctx, TraceIdKey, "no-trace-id")
	}

	slog.LogAttrs(
		ctx,
		logLevel,
		message,
		append(args, slog.String("traceId", ctx.Value(TraceIdKey).(string)))...,
	)
}
