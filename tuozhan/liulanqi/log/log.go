package log

import (
	"os"

	"github.com/gookit/slog"
)

var std = &slog.SugaredLogger{}

func init() {
	std = newStdLogger(slog.NoticeLevel)
}

const template = "[{{level}}] {{message}} {{data}} {{extra}}\n"

// newStdLogger is a new std logger
func newStdLogger(level slog.Level) *slog.SugaredLogger {
	return slog.NewSugaredLogger(os.Stdout, level).Config(func(sl *slog.SugaredLogger) {
		sl.SetName("stdLogger")
		sl.ReportCaller = true
		sl.CallerSkip = 7
		// auto enable console color
		sl.Formatter.(*slog.TextFormatter).EnableColor = false
		sl.Formatter.(*slog.TextFormatter).SetTemplate(template)
	})
}

// Infof logs a message at level Info
func Infof(format string, args ...interface{}) {
	std.Logf(slog.InfoLevel, format, args...)
}

// Notice logs a message at level Notice
func Notice(args ...interface{}) {
	std.Log(slog.NoticeLevel, args...)
}

// Noticef logs a message at level Notice
func Noticef(format string, args ...interface{}) {
	std.Logf(slog.NoticeLevel, format, args...)
}

// Warn logs a message at level Warn
func Warn(args ...interface{}) {
	std.Log(slog.WarnLevel, args...)
}

// Warnf logs a message at level Warn

// Error logs a message at level Error
func Error(args ...interface{}) {
	std.Log(slog.ErrorLevel, args...)
}

// Errorf logs a message at level Error
func Errorf(format string, args ...interface{}) {
	std.Logf(slog.ErrorLevel, format, args...)
}

// Fatal logs a message at level Fatal
func Fatal(args ...interface{}) {
	std.Log(slog.FatalLevel, args...)
}

// Fatalf logs a message at level Fatal
func Fatalf(format string, args ...interface{}) {
	std.Logf(slog.FatalLevel, format, args...)
}
