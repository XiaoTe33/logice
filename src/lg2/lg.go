package lg2

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var levelMap = map[logrus.Level]string{
	logrus.TraceLevel: "[TRA]",
	logrus.DebugLevel: "[DEB]",
	logrus.InfoLevel:  "[INF]",
	logrus.WarnLevel:  "[WAR]",
	logrus.ErrorLevel: "[ERR]",
	logrus.PanicLevel: "[PAN]",
	logrus.FatalLevel: "[FAT]",
}

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

type Formatters struct {
	Info  F
	Debug F
	Warn  F
	Error F
}

type Logger struct {
	outputs []Outputs
	Formatters
}

type F struct {
	jump     bool
	funcName bool
	color    int
	bgColor  int
}

type logFormatter struct {
}

func (logger *Logger) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = Gray
	case logrus.WarnLevel:
		levelColor = Yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = Red
	default:
		levelColor = Blue
	}
	b := &bytes.Buffer{}
	if entry.Buffer != nil {
		b = entry.Buffer
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		//funcVal := entry.Caller.Function

		dir, err2 := os.Getwd()
		if err2 != nil {
			panic(err2)
		}
		Path, err := filepath.Rel(dir, path.Dir(entry.Caller.File))
		if err != nil {
			panic(err)
		}
		_ = fmt.Sprintf(".\\%s:%d", Path+"\\"+path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[1;4%dm%s\x1b[0m ", timestamp, levelColor, levelMap[entry.Level])
		if logger.Info.jump {
			fmt.Fprintf(b, "%-40s|", GetJump(9))
		}
		if logger.Info.funcName {
			split := strings.Split(getFuncName(9), ".")
			before := strings.Join(split[:len(split)-1], ".")
			funName := split[len(split)-1]
			fmt.Fprintf(b, " %-40s| ", before+".\x1b[36m"+funName+"\x1b[0m")
		}
		fmt.Fprintf(b, "\"%s\"\n", entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm%s\x1b[0m %s\n", timestamp, levelColor, levelMap[entry.Level], entry.Message)
	}
	return b.Bytes(), nil
}

func WithOutputStd(filepath string, levels ...logrus.Level) OptionFunc {
	return func(logger *Logger) {
		logger.outputs = append(logger.outputs, Outputs{
			Logger:    logrus.Logger{},
			Filepath:  filepath,
			Formatter: logger,
			Levels:    levels,
		})
	}
}
func WithOutputFile(filepath string, levels ...logrus.Level) OptionFunc {
	return func(logger *Logger) {
		logger.outputs = append(logger.outputs, Outputs{
			Logger:    logrus.Logger{},
			Filepath:  filepath,
			Formatter: logger,
			Levels:    levels,
		})
	}
}

func WithFunc(levels ...logrus.Level) OptionFunc {
	return func(logger *Logger) {

	}
}

func (logger *Logger) SetColor(level logrus.Level, color int) {
	switch level {
	case logrus.InfoLevel:
		logger.Info.color = color
	case logrus.DebugLevel:
		logger.Debug.color = color
	case logrus.WarnLevel:
		logger.Warn.color = color
	case logrus.ErrorLevel:
		logger.Error.color = color
	}
}

func (logger *Logger) SetBgColor(level logrus.Level, color int) {
	switch level {
	case logrus.InfoLevel:
		logger.Info.bgColor = color
	case logrus.DebugLevel:
		logger.Debug.bgColor = color
	case logrus.WarnLevel:
		logger.Warn.bgColor = color
	case logrus.ErrorLevel:
		logger.Error.bgColor = color
	}
}

func (logger *Logger) SetFormatter(level logrus.Level, formatter *logrus.Formatter) {

}

func AddOutput(filepath string, levels ...logrus.Level) OptionFunc {
	return func(logger *Logger) {
		logger.outputs = append(logger.outputs, Outputs{
			Logger:    logrus.Logger{},
			Filepath:  filepath,
			Formatter: nil,
			Levels:    levels,
		})
	}
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
	Formatter logrus.Formatter
	Levels    []logrus.Level
}

const (
	InfoLevel  = logrus.InfoLevel
	DebugLevel = logrus.DebugLevel
	WarnLevel  = logrus.WarnLevel
	ErrorLevel = logrus.ErrorLevel
)

func getFuncName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		panic("runtime.Caller() failed")
	}
	funcName := runtime.FuncForPC(pc).Name()
	return funcName
}

func GetJump(skip int) string {
	_, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		panic("runtime.Caller() failed")
	}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	Path, err := filepath.Rel(dir, file)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_ = path.Base(file) // Base函数返回路径的最后一个元素
	return fmt.Sprintf(".\\%s:%d", Path, lineNo)
}
