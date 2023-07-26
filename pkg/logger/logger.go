package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"path"
)

const (
	Black = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Yan
	Gray
)

type Logger struct {
}

func Init(option *Options) {
	//设置日志级别
	logrus.SetLevel(option.Level)

	//是否打印输出信息
	logrus.SetReportCaller(option.EnableCaller)

	//设置输出位置
	logrus.SetOutput(io.MultiWriter(option.OutputPaths...))

	//添加hook 区分错误信息
	logrus.AddHook(&MyHook{Write: option.ErrorOutputPaths})

	if option.Format == "json" {
		logrus.SetFormatter(&MyForJson{EnableColor: option.EnableColor, EnableCaller: option.EnableCaller})
	} else {
		logrus.SetFormatter(&MyForConsole{EnableColor: option.EnableColor, EnableCaller: option.EnableCaller})
	}

}

type MyForJson struct {
	EnableColor      bool        `json:"enable-color" mapstructure:"enable-color"`
	EnableCaller     bool        `json:"enable-caller" mapstructure:"enable-caller"`
	OutputPaths      []io.Writer `json:"output-paths" mapstructure:"output-paths"`
	ErrorOutputPaths []io.Writer `json:"error-output-paths" mapstructure:"error-output-paths"`
}

func (m *MyForJson) Format(entry *logrus.Entry) ([]byte, error) {
	var color int
	switch entry.Level {
	case logrus.ErrorLevel:
		color = Red
	case logrus.WarnLevel:
		color = Yellow
	case logrus.InfoLevel:
		color = Blue
	case logrus.DebugLevel:
		color = Green
	default:
		color = Purple
	}
	var b *bytes.Buffer
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}
	//时间格式化
	formatTime := entry.Time.Format("2006-01-02 15:04:06")
	entry.Data["timestamp"] = formatTime
	m.InEntry(entry)
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		//fmt.Println(funcVal)
		fileVal := fmt.Sprintf("%s:[%s:%d]", path.Base(entry.Caller.File), funcVal, entry.Caller.Line)
		entry.Data["caller"] = fileVal
	}

	//自定义文件路径
	resultBytes, _ := json.Marshal(entry.Data)
	if m.EnableColor {
		fmt.Fprintf(b, "\033[3%dm%v\033[0m \n", color, string(resultBytes))
	} else {
		fmt.Fprintf(b, "%s \n", string(resultBytes))
	}

	return b.Bytes(), nil
}

func (m *MyForJson) InEntry(entry *logrus.Entry) {
	entry.Data["level"] = entry.Level
	entry.Data["message"] = entry.Message

}

type MyForConsole struct {
	EnableColor      bool        `json:"enable-color" mapstructure:"enable-color"`
	EnableCaller     bool        `json:"enable-caller" mapstructure:"enable-caller"`
	OutputPaths      []io.Writer `json:"output-paths" mapstructure:"output-paths"`
	ErrorOutputPaths []io.Writer `json:"error-output-paths" mapstructure:"error-output-paths"`
	MyForJson
}

func (m *MyForConsole) Format(entry *logrus.Entry) ([]byte, error) {
	//TODO implement me
	var color int
	switch entry.Level {
	case logrus.ErrorLevel:
		color = Red
	case logrus.WarnLevel:
		color = Yellow
	case logrus.InfoLevel:
		color = Blue
	case logrus.DebugLevel:
		color = Green
	default:
		color = Purple
	}
	var b *bytes.Buffer
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}
	//时间格式化
	formatTime := entry.Time.Format("2006-01-02 15:04:06")

	//是否打印颜色
	if m.EnableColor {
		fmt.Fprintf(b, "\033[3%dm[%s]\033[0m", color, entry.Level)
	} else {
		fmt.Fprintf(b, "[%s]", entry.Level)
	}

	fmt.Fprintf(b, " %s", formatTime)
	//是否添加调用信息
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		//fmt.Println(funcVal)
		fileVal := fmt.Sprintf("%s:[%s:%d]", path.Base(entry.Caller.File), funcVal, entry.Caller.Line)
		fmt.Fprintf(b, " %s ", fileVal)
	}

	//添加时间 +日志内容
	fmt.Fprintf(b, " %s", entry.Message)
	if len(entry.Data) != 0 {
		bytess, _ := json.Marshal(entry.Data)
		fmt.Fprintf(b, " %v \n", string(bytess))
	} else {
		fmt.Fprintf(b, " \n")
	}

	return b.Bytes(), nil
}

type MyHook struct {
	Write []io.Writer
}

// Levels  这里设置error 错误级别的日志指定输出到ErrorOutputPaths 中
func (m MyHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (m MyHook) Fire(entry *logrus.Entry) error {
	//TODO implement me
	//bytes, _ := m.MyForConsole.Format(entry)
	//m.Write.Write(bytes)
	line, _ := entry.String()
	for _, w := range m.Write {
		w.Write([]byte(line))
	}
	return nil
}
