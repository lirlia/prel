package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
)

type loggerKey struct{}

type builder struct {
	output io.Writer
	opts   slog.HandlerOptions
}

type Builder interface {
	Build() *slog.Logger
	WithOutput(w io.Writer) Builder
	WithOptions(opts slog.HandlerOptions) Builder
}

var defaultOutput io.Writer = os.Stdout

func NewBuilder() Builder {
	return &builder{
		output: defaultOutput,
	}
}

func (b *builder) Build() *slog.Logger {
	return slog.New(slog.NewJSONHandler(b.output, &b.opts))
}

func (b *builder) WithOutput(w io.Writer) Builder {
	b.output = w
	return b
}

func (b *builder) WithOptions(opts slog.HandlerOptions) Builder {
	b.opts = opts
	return b
}

func Set(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func Get(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(loggerKey{}).(*slog.Logger)
	if !ok {
		fmt.Println("logger is nil")
		return slog.Default()
	}
	return logger
}

func ReplaceAttr(groups []string, a slog.Attr) slog.Attr {
	switch a.Key {
	case slog.LevelKey:
		// Cloud Logging: level -> severity
		severity, ok := a.Value.Any().(slog.Level)
		if !ok {
			return a
		}
		return slog.String("severity", severity.String())
	case slog.MessageKey:
		// Cloud Logging: msg -> message
		return slog.Attr{Key: "message", Value: a.Value}
	case slog.SourceKey:
		source, ok := a.Value.Any().(*slog.Source)
		if !ok {
			return a
		}

		return slog.String("caller", fmt.Sprintf("%s:%d", source.File, source.Line))
	}
	return a
}
