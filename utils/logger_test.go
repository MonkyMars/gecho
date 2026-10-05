package utils

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestLevelStringAndParseLevel(t *testing.T) {
	for level, want := range map[Level]string{
		LevelDebug: "DEBUG",
		LevelInfo:  "INFO",
		LevelWarn:  "WARN",
		LevelError: "ERROR",
		LevelFatal: "FATAL",
		Level(99):  "UNKNOWN",
	} {
		if got := level.String(); got != want {
			t.Errorf("Level(%d).String() = %q, want %q", level, got, want)
		}
	}

	for input, want := range map[string]Level{
		"debug": LevelDebug, "INFO": LevelInfo, "warn": LevelWarn,
		"WARNING": LevelWarn, "error": LevelError, "fatal": LevelFatal,
		"unknown": LevelInfo, "": LevelInfo,
	} {
		if got := ParseLevel(input); got != want {
			t.Errorf("ParseLevel(%q) = %v, want %v", input, got, want)
		}
	}
}

func TestConfigOptions(t *testing.T) {
	var output, errorOutput bytes.Buffer
	config := NewConfig(
		WithLogLevel(LevelDebug),
		WithLogFormat(FormatJSON),
		WithColorize(true),
		WithShowCaller(false),
		WithTimeFormat(timeFormatForTests),
		WithOutput(&output),
		WithErrorOutput(&errorOutput),
		WithDefaultCallerSkip(7),
	)

	if config.Level != LevelDebug || config.Format != FormatJSON ||
		!config.Colorize || config.ShowCaller || config.CallerSkip != 7 ||
		config.TimeFormat != timeFormatForTests || config.Output != &output ||
		config.ErrorOutput != &errorOutput {
		t.Fatalf("options were not applied: %#v", config)
	}
}

const timeFormatForTests = "15:04:05"

func TestLoggerLevelsAndOutputs(t *testing.T) {
	var output, errorOutput bytes.Buffer
	logger := NewLogger(NewConfig(
		WithLogLevel(LevelInfo),
		WithLogFormat(FormatJSON),
		WithShowCaller(false),
		WithOutput(&output),
		WithErrorOutput(&errorOutput),
		WithTimeFormat(timeFormatForTests),
	))

	logger.Debug("hidden")
	if output.Len() != 0 {
		t.Fatalf("debug message was written below configured level: %q", output.String())
	}
	logger.Info("visible", Field("request_id", "abc"))
	logger.Warn("warning")
	logger.Error("failure")

	if !strings.Contains(output.String(), `"level":"INFO"`) ||
		!strings.Contains(output.String(), `"request_id":"abc"`) ||
		!strings.Contains(output.String(), `"level":"WARN"`) {
		t.Fatalf("unexpected info/warn output: %q", output.String())
	}
	if errorOutput.Len() == 0 || !strings.Contains(errorOutput.String(), `"level":"ERROR"`) {
		t.Fatalf("error was not routed to error output: %q", errorOutput.String())
	}
}

func TestLoggerFormatsAndFields(t *testing.T) {
	tests := []struct {
		name   string
		format Format
		check  func(*testing.T, string)
	}{
		{"text", FormatText, func(t *testing.T, got string) {
			if !strings.Contains(got, "INFO  hello") || !strings.Contains(got, "id=42") {
				t.Errorf("unexpected text output: %q", got)
			}
		}},
		{"json", FormatJSON, func(t *testing.T, got string) {
			var entry struct {
				Level  string         `json:"level"`
				Msg    string         `json:"message"`
				Fields map[string]any `json:"fields"`
			}
			if err := json.Unmarshal([]byte(got), &entry); err != nil {
				t.Fatal(err)
			}
			if entry.Level != "INFO" || entry.Msg != "hello" || entry.Fields["id"] != float64(42) {
				t.Errorf("unexpected JSON entry: %#v", entry)
			}
		}},
		{"pretty", FormatPretty, func(t *testing.T, got string) {
			if !strings.Contains(got, "INFO") || !strings.Contains(got, "hello") ||
				!strings.Contains(got, "(id=42)") {
				t.Errorf("unexpected pretty output: %q", got)
			}
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			logger := NewLogger(NewConfig(
				WithLogFormat(tt.format),
				WithShowCaller(false),
				WithColorize(false),
				WithOutput(&output),
				WithErrorOutput(&output),
			))
			logger.Info("hello", Field("id", 42))
			tt.check(t, output.String())
		})
	}
}

func TestLoggerPersistentFieldsAreCopied(t *testing.T) {
	var output bytes.Buffer
	logger := NewLogger(NewConfig(WithLogFormat(FormatJSON), WithShowCaller(false), WithOutput(&output)))
	child := logger.WithFields(map[string]any{"request": "one"}).WithField("user", 2)

	logger.Info("parent")
	if strings.Contains(output.String(), "request") {
		t.Fatal("parent logger inherited child fields")
	}
	output.Reset()
	child.Info("child")
	if !strings.Contains(output.String(), `"request":"one"`) ||
		!strings.Contains(output.String(), `"user":2`) {
		t.Fatalf("child fields missing: %q", output.String())
	}
}

func TestParseArgs(t *testing.T) {
	field := Field("key", "value")
	msg, opts := parseArgs("message", field, "ignored")
	if msg != "message" || len(opts) != 1 {
		t.Fatalf("parseArgs with message = %q, %d options", msg, len(opts))
	}
	msg, opts = parseArgs(field)
	if msg != "" || len(opts) != 1 {
		t.Fatalf("parseArgs without message = %q, %d options", msg, len(opts))
	}
	msg, opts = parseArgs()
	if msg != "" || len(opts) != 0 {
		t.Fatalf("parseArgs empty = %q, %d options", msg, len(opts))
	}
}

func TestLoggerSettersAndCallerOption(t *testing.T) {
	var output bytes.Buffer
	logger := NewLogger(NewConfig(
		WithLogFormat(FormatText),
		WithShowCaller(false),
		WithOutput(&output),
		WithErrorOutput(&output),
	))
	logger.SetLevel(LevelError)
	logger.Info("hidden")
	if output.Len() != 0 {
		t.Fatal("SetLevel did not suppress info")
	}

	logger.SetFormat(FormatJSON)
	logger.Error("shown", WithCallerSkip(0))
	if !strings.Contains(output.String(), `"level":"ERROR"`) {
		t.Fatalf("SetFormat did not select JSON: %q", output.String())
	}
}

func TestNewDefaultLogger(t *testing.T) {
	if logger := NewDefaultLogger(); logger == nil {
		t.Fatal("NewDefaultLogger returned nil")
	}
}
