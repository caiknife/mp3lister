package logger

import (
	"os"

	"github.com/golang-module/carbon/v2"
	"github.com/sirupsen/logrus"
)

var (
	ConsoleLogger *logrus.Logger
)

func init() {
	ConsoleLogger = NewConsoleLogger()
}

func NewConsoleLogger() *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
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
	})
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.TraceLevel)
	return l
}
