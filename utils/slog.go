package utils

import (
	"context"
	"log/slog"
	"runtime"
	"strconv"
	"strings"
)

type slogHandler struct {
	logger *Logger
	attrs  []slog.Attr
	group  string
}

func (h *slogHandler) Enabled(_ context.Context, level slog.Level) bool {
	h.logger.mu.Lock()
	defer h.logger.mu.Unlock()
	return level >= slogLevel(h.logger.config.Level)
}

func (h *slogHandler) Handle(_ context.Context, record slog.Record) error {
	h.logger.mu.Lock()
	defer h.logger.mu.Unlock()

	level := levelFromSlog(record.Level)
	if level < h.logger.config.Level {
		return nil
	}

	fields := make([]fieldPair, 0, len(h.attrs)+record.NumAttrs())
	for _, attr := range h.attrs {
		appendSlogAttr(&fields, attr, h.group)
	}
	record.Attrs(func(attr slog.Attr) bool {
		appendSlogAttr(&fields, attr, h.group)
		return true
	})

	entry := entry{
		Timestamp: record.Time.Format(h.logger.config.TimeFormat),
		Level:     level.String(),
		Message:   record.Message,
		Fields:    make([]fieldPair, 0, len(h.logger.fields)+len(fields)),
	}
	for key, value := range h.logger.fields {
		entry.Fields = append(entry.Fields, fieldPair{Key: key, Value: value})
	}
	entry.Fields = append(entry.Fields, fields...)

	if h.logger.config.ShowCaller && record.PC != 0 {
		if fn := runtime.FuncForPC(record.PC); fn != nil {
			file, line := fn.FileLine(record.PC)
			parts := strings.Split(file, "/")
			entry.Caller = parts[len(parts)-1] + ":" + formatLine(line)
		}
	}

	output := h.logger.config.Output
	if level >= LevelError && h.logger.config.ErrorOutput != nil {
		output = h.logger.config.ErrorOutput
	}
	switch h.logger.config.Format {
	case FormatJSON:
		h.logger.writeJSON(output, entry)
	case FormatPretty:
		h.logger.writePretty(output, level, entry)
	default:
		h.logger.writeText(output, level, entry)
	}
	return nil
}

func (h *slogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	combined := make([]slog.Attr, 0, len(h.attrs)+len(attrs))
	combined = append(combined, h.attrs...)
	combined = append(combined, attrs...)
	return &slogHandler{logger: h.logger, attrs: combined, group: h.group}
}

func (h *slogHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	group := name
	if h.group != "" {
		group = h.group + "." + name
	}
	return &slogHandler{logger: h.logger, attrs: h.attrs, group: group}
}

func appendSlogAttr(fields *[]fieldPair, attr slog.Attr, group string) {
	if attr.Key == "" {
		return
	}
	attr.Value = attr.Value.Resolve()
	key := attr.Key
	if group != "" {
		key = group + "." + key
	}
	if attr.Value.Kind() == slog.KindGroup {
		for _, child := range attr.Value.Group() {
			appendSlogAttr(fields, child, key)
		}
		return
	}
	*fields = append(*fields, fieldPair{Key: key, Value: attr.Value.Any()})
}

func slogLevel(level Level) slog.Level {
	switch level {
	case LevelDebug:
		return slog.LevelDebug
	case LevelWarn:
		return slog.LevelWarn
	case LevelError, LevelFatal:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func levelFromSlog(level slog.Level) Level {
	switch {
	case level <= slog.LevelDebug:
		return LevelDebug
	case level < slog.LevelWarn:
		return LevelInfo
	case level < slog.LevelError:
		return LevelWarn
	default:
		return LevelError
	}
}

func formatLine(line int) string {
	return strconv.Itoa(line)
}
