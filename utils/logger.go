package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Level represents the severity of a log message
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// String returns the string representation of the log level
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// ParseLevel converts a string to a Level
func ParseLevel(s string) Level {
	switch strings.ToLower(s) {
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "warn", "warning":
		return LevelWarn
	case "error":
		return LevelError
	case "fatal":
		return LevelFatal
	default:
		return LevelInfo
	}
}

// Format defines how logs should be formatted
type Format int

const (
	FormatText Format = iota
	FormatJSON
)

// Config contains logger configuration
type Config struct {
	Level       Level
	Format      Format
	Output      io.Writer
	ErrorOutput io.Writer
	Colorize    bool
	ShowCaller  bool
	CallerSkip  int
	TimeFormat  string
}

// DefaultConfig returns a logger config with sensible defaults
func DefaultConfig() Config {
	return Config{
		Level:       LevelInfo,
		Format:      FormatText,
		Output:      os.Stdout,
		ErrorOutput: os.Stderr,
		Colorize:    isTerminal(os.Stdout),
		ShowCaller:  true,
		CallerSkip:  2,
		TimeFormat:  "2006-01-02 15:04:05.000",
	}
}

type LoggerOptions func(*Config)

func NewConfig(options ...LoggerOptions) Config {
	config := DefaultConfig()
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithLogLevel(level Level) LoggerOptions {
	return func(c *Config) {
		c.Level = level
	}
}

func WithLogFormat(format Format) LoggerOptions {
	return func(c *Config) {
		c.Format = format
	}
}

func WithColorize(colorize bool) LoggerOptions {
	return func(c *Config) {
		c.Colorize = colorize
	}
}

func WithShowCaller(showCaller bool) LoggerOptions {
	return func(c *Config) {
		c.ShowCaller = showCaller
	}
}

func WithTimeFormat(timeFormat string) LoggerOptions {
	return func(c *Config) {
		c.TimeFormat = timeFormat
	}
}

func WithOutput(output io.Writer) LoggerOptions {
	return func(c *Config) {
		c.Output = output
	}
}

func WithErrorOutput(errorOutput io.Writer) LoggerOptions {
	return func(c *Config) {
		c.ErrorOutput = errorOutput
	}
}

func WithDefaultCallerSkip(callerSkip int) LoggerOptions {
	return func(c *Config) {
		c.CallerSkip = callerSkip
	}
}

// Logger is a structured, thread-safe logger
type Logger struct {
	mu     sync.Mutex
	config Config
	fields map[string]any
}

// New creates a new logger with the given configuration
func NewLogger(config Config) *Logger {
	return &Logger{
		config: config,
		fields: make(map[string]any),
	}
}

// NewDefaultLogger creates a new logger with default configuration
func NewDefaultLogger() *Logger {
	return NewLogger(DefaultConfig())
}

// WithField returns a new logger with an additional field
func (l *Logger) WithField(key string, value any) *Logger {
	return l.WithFields(map[string]any{key: value})
}

// WithFields returns a new logger with additional fields
func (l *Logger) WithFields(fields map[string]any) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()

	newFields := maps.Clone(l.fields)
	maps.Copy(newFields, fields)

	return &Logger{
		config: l.config,
		fields: newFields,
	}
}

// SetLevel sets the minimum log level
func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config.Level = level
}

// SetFormat sets the output format
func (l *Logger) SetFormat(format Format) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config.Format = format
}

// entry represents a single log entry
type entry struct {
	Timestamp string         `json:"timestamp"`
	Level     string         `json:"level"`
	Message   string         `json:"message"`
	Caller    string         `json:"caller,omitempty"`
	Fields    map[string]any `json:"fields,omitempty"`
}

type Option func(*entryOptions)

type entryOptions struct {
	callerSkip *int
	fields     map[string]any
}

func WithCallerSkip(skip int) Option {
	return func(o *entryOptions) {
		o.callerSkip = &skip
	}
}

func (l *Logger) log(level Level, msg string, opts []Option) {
	if level < l.config.Level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	o := entryOptions{
		fields: make(map[string]any),
	}

	for _, opt := range opts {
		opt(&o)
	}

	e := entry{
		Timestamp: time.Now().Format(l.config.TimeFormat),
		Level:     level.String(),
		Message:   msg,
		Fields:    make(map[string]any),
	}

	// Persistent fields
	maps.Copy(e.Fields, l.fields)

	// Option fields
	maps.Copy(e.Fields, o.fields)

	// Caller handling
	callerSkip := l.config.CallerSkip
	if o.callerSkip != nil {
		callerSkip = *o.callerSkip
	}

	if l.config.ShowCaller {
		if _, file, line, ok := runtime.Caller(callerSkip); ok {
			parts := strings.Split(file, "/")
			e.Caller = fmt.Sprintf("%s:%d", parts[len(parts)-1], line)
		}
	}

	// Output selection
	output := l.config.Output
	if level >= LevelError && l.config.ErrorOutput != nil {
		output = l.config.ErrorOutput
	}

	// Write
	if l.config.Format == FormatJSON {
		l.writeJSON(output, e)
	} else {
		l.writeText(output, level, e)
	}

	if level == LevelFatal {
		os.Exit(1)
	}
}

var levelColors = map[Level]string{
	LevelDebug: "\033[36m", // Cyan
	LevelInfo:  "\033[32m", // Green
	LevelWarn:  "\033[33m", // Yellow
	LevelError: "\033[31m", // Red
	LevelFatal: "\033[35m", // Magenta
}

const colorReset = "\033[0m"

// writeText writes the entry in human-readable text format
func (l *Logger) writeText(w io.Writer, level Level, e entry) {
	var sb strings.Builder

	// Timestamp
	sb.WriteString(e.Timestamp)
	sb.WriteString(" ")

	// Level with optional color
	if l.config.Colorize {
		sb.WriteString(levelColors[level])
	}
	sb.WriteString(fmt.Sprintf("%-5s", e.Level))
	if l.config.Colorize {
		sb.WriteString(colorReset)
	}
	sb.WriteString(" ")

	// Caller
	if e.Caller != "" {
		sb.WriteString("[")
		sb.WriteString(e.Caller)
		sb.WriteString("] ")
	}

	// Message
	sb.WriteString(e.Message)

	// Fields
	if len(e.Fields) > 0 {
		sb.WriteString(" {")
		first := true
		for k, v := range e.Fields {
			if !first {
				sb.WriteString(", ")
			}
			sb.WriteString(k)
			sb.WriteString("=")
			sb.WriteString(fmt.Sprint(v))
			first = false
		}
		sb.WriteString("}")
	}

	fmt.Fprintln(w, sb.String())
}

// writeJSON writes the entry in JSON format
func (l *Logger) writeJSON(w io.Writer, e entry) {
	data, _ := json.Marshal(e)
	fmt.Fprintln(w, string(data))
}

// Debug logs a debug level message
func (l *Logger) Debug(args ...any) {
	msg, opts := parseArgs(args...)
	l.log(LevelDebug, msg, opts)
}

// Info logs an info level message
func (l *Logger) Info(args ...any) {
	msg, opts := parseArgs(args...)
	l.log(LevelInfo, msg, opts)
}

// Warn logs a warning level message
func (l *Logger) Warn(args ...any) {
	msg, opts := parseArgs(args...)
	l.log(LevelWarn, msg, opts)
}

// Error logs an error level message
func (l *Logger) Error(args ...any) {
	msg, opts := parseArgs(args...)
	l.log(LevelError, msg, opts)
}

// Fatal logs a fatal level message and exits the program
func (l *Logger) Fatal(args ...any) {
	msg, opts := parseArgs(args...)
	l.log(LevelFatal, msg, opts)
}

// isTerminal checks if the writer is a terminal
func isTerminal(w io.Writer) bool {
	if f, ok := w.(*os.File); ok {
		stat, err := f.Stat()
		if err != nil {
			return false
		}
		return (stat.Mode() & os.ModeCharDevice) != 0
	}
	return false
}

func Field(key string, value any) Option {
	return func(o *entryOptions) {
		if o.fields == nil {
			o.fields = make(map[string]any)
		}
		o.fields[key] = value
	}
}

func parseArgs(args ...any) (string, []Option) {
	var msg string
	opts := make([]Option, 0, len(args))

	if len(args) == 0 {
		return "", opts
	}

	// First argument may be a message string
	if s, ok := args[0].(string); ok {
		msg = s
		for _, a := range args[1:] {
			if opt, ok := a.(Option); ok {
				opts = append(opts, opt)
			}
		}
	} else {
		for _, a := range args {
			if opt, ok := a.(Option); ok {
				opts = append(opts, opt)
			}
		}
	}

	return msg, opts
}
