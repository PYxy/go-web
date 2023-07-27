package main

import (
	"fmt"
	_ "github.com/PYxy/go-web/internal/pkg/logger"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
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
