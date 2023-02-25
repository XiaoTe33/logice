package lg2

import (
	"github.com/sirupsen/logrus"
)

const (
	Black = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan //青色
	Gray
)

const (
	StdOut  = 0
	FileOut = 1
)

type Color struct {
	cInfo    int
	cDebug   int
	cWarn    int
	cError   int
	bgcInfo  int
	bgcDebug int
	bgcWarn  int
	bgcError int
}

type Logger struct {
	outputs []Outputs
	Color
}

func WithOutputStd(filepath string) OptionFunc {
	return func(logger *Logger) {

	}
}
func WithOutputFile(level logrus.Level, filepath string) OptionFunc {
	return func(logger *Logger) {
		logger.outputs = append(logger.outputs, Outputs{
			Logger:    logrus.Logger{},
			Filepath:  filepath,
			Formatter: nil,
			Levels:    nil,
		})
	}
}

func With() {

}

func (logger *Logger) SetColor(level logrus.Level, color int) {
	switch level {
	case logrus.InfoLevel:
		logger.cInfo = color
	case logrus.DebugLevel:
		logger.cDebug = color
	case logrus.WarnLevel:
		logger.cWarn = color
	case logrus.ErrorLevel:
		logger.cError = color
	}
}

func (logger *Logger) SetBgColor(level logrus.Level, color int) {
	switch level {
	case logrus.InfoLevel:
		logger.bgcInfo = color
	case logrus.DebugLevel:
		logger.bgcDebug = color
	case logrus.WarnLevel:
		logger.bgcWarn = color
	case logrus.ErrorLevel:
		logger.bgcError = color
	}
}

func (logger *Logger) SetFormatter(level logrus.Level, formatter *logrus.Formatter) {

}

func New(options ...OptionFunc) *Logger {
	logger := &Logger{}
	for _, option := range options {
		option(logger)
	}
	return logger
}

type OptionFunc func(*Logger)

type Option struct {
}

type Outputs struct {
	Logger    logrus.Logger
	Filepath  string
	Formatter *logrus.Formatter
	Levels    []logrus.Level
}

const (
	InfoLevel  = logrus.InfoLevel
	DebugLevel = logrus.DebugLevel
	WarnLevel  = logrus.WarnLevel
	ErrorLevel = logrus.ErrorLevel
)
