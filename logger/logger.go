package logger

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/matteoaricci/jot-api/version"
)

const modulePath = "github.com/matteoaricci/jot-api/"

// trimSource rewrites the source attr's file to a repo-relative path and its
// function to drop the module prefix, so logs are stable and searchable across
// environments (container/Lambda paths differ from local absolute paths).
func trimSource(groups []string, a slog.Attr) slog.Attr {
	if a.Key != slog.SourceKey {
		return a
	}
	src, ok := a.Value.Any().(*slog.Source)
	if !ok {
		return a
	}
	if i := strings.LastIndex(src.File, modulePath); i >= 0 {
		src.File = src.File[i+len(modulePath):]
	}
	src.Function = strings.TrimPrefix(src.Function, modulePath)
	return a
}

type ctxKey struct{}

// WithAttrs returns a child context carrying attributes that the context-aware
// handler adds to every record logged with that context.
func WithAttrs(ctx context.Context, attrs ...slog.Attr) context.Context {
	existing, _ := ctx.Value(ctxKey{}).([]slog.Attr)
	combined := make([]slog.Attr, 0, len(existing)+len(attrs))
	combined = append(combined, existing...)
	combined = append(combined, attrs...)
	return context.WithValue(ctx, ctxKey{}, combined)
}

type contextHandler struct {
	slog.Handler
}

func (h contextHandler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(ctxKey{}).([]slog.Attr); ok {
		r.AddAttrs(attrs...)
	}
	return h.Handler.Handle(ctx, r)
}

func (h contextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return contextHandler{h.Handler.WithAttrs(attrs)}
}

func (h contextHandler) WithGroup(name string) slog.Handler {
	return contextHandler{h.Handler.WithGroup(name)}
}

// Init installs a JSON slog logger (with source) as the default, tagged with
// app.name and wrapped to inject request-scoped context attributes.
func Init() {
	base := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		AddSource:   true,
		ReplaceAttr: trimSource,
	})

	info := version.GetInfo()
	h := slog.New(contextHandler{base}).
		With(slog.Group("app",
			slog.String("name", "jot-api"),
			slog.String("version", info.Version),
			slog.String("commit", info.GitCommit),
		))

	slog.SetDefault(h)
}
