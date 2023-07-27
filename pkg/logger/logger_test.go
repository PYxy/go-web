package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	fmt.Println("文件打开异常:", err)
	defer file.Close()
	file1, err := os.OpenFile("err.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	fmt.Println("文件打开异常:", err)
	defer file1.Close()

	//option := &logger.Options{
	//	Level:            logrus.DebugLevel,
	//	Format:           "console", //json console
	//	EnableColor:      true,
	//	EnableCaller:     true,
	//	OutputPaths:      []io.Writer{os.Stdout, file},
	//	ErrorOutputPaths: []io.Writer{file1},
	//}
	//logger.Init(option)
	logrus.Debug("你好")
	logrus.Info("nihao2")
	logrus.WithField("name", "ljy").Info("lalala")
	logrus.WithField("name", "ljy2").Error("lalala")

}

func TestLogger2(t *testing.T) {
	option := &Options{
		Level:            logrus.DebugLevel,
		Format:           "json",
		EnableColor:      true,
		EnableCaller:     false,
		OutputPaths:      nil,
		ErrorOutputPaths: nil,
	}
	Init(option)
	logrus.Debug("你好")
}
