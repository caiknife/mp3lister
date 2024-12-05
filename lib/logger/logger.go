package logger

import (
	"os"

	"github.com/dromara/carbon/v2"
	"github.com/sirupsen/logrus"
)

var (
	ConsoleLogger *logrus.Logger
)

func init() {
	ConsoleLogger = NewConsoleLogger()
}

func NewConsoleLogger(opts ...ConsoleLoggerFormatterOption) *logrus.Logger {
	l := logrus.New()
	f := &logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              true,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           carbon.ISO8601MicroLayout,
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	}
	for _, opt := range opts {
		opt(f)
	}
	l.SetFormatter(f)
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.TraceLevel)
	return l
}

type ConsoleLoggerFormatterOption func(formatter *logrus.TextFormatter)

func TimestampFormat(format string) ConsoleLoggerFormatterOption {
	return func(formatter *logrus.TextFormatter) {
		formatter.TimestampFormat = format
	}
}

func FullTimestamp(flag bool) ConsoleLoggerFormatterOption {
	return func(formatter *logrus.TextFormatter) {
		formatter.FullTimestamp = flag
	}
}

func DisableTimestamp(flag bool) ConsoleLoggerFormatterOption {
	return func(formatter *logrus.TextFormatter) {
		formatter.DisableTimestamp = flag
	}
}

func EnvironmentOverrideColors(flag bool) ConsoleLoggerFormatterOption {
	return func(formatter *logrus.TextFormatter) {
		formatter.ForceColors = flag
	}
}

func DisableQuote(flag bool) ConsoleLoggerFormatterOption {
	return func(formatter *logrus.TextFormatter) {
		formatter.DisableQuote = flag
	}
}

func ForceQuote(flag bool) ConsoleLoggerFormatterOption {
	return func(formatter *logrus.TextFormatter) {
		formatter.ForceQuote = flag
	}
}

func DisableColors(flag bool) ConsoleLoggerFormatterOption {
	return func(formatter *logrus.TextFormatter) {
		formatter.DisableColors = flag
	}
}

func ForceColors(flag bool) ConsoleLoggerFormatterOption {
	return func(formatter *logrus.TextFormatter) {
		formatter.ForceColors = flag
	}
}
