package main

import (
	"Base-Pro/pro_modle/internal/pkg/logger"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	fmt.Println("文件打开异常:", err)
	defer file.Close()
	file1, err := os.OpenFile("err.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	fmt.Println("文件打开异常:", err)
	defer file1.Close()

	file2, err := os.OpenFile("err2.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	fmt.Println("文件打开异常:", err)
	defer file2.Close()

	option := &logger.Options{
		Level:            logrus.DebugLevel,
		Format:           "console",
		EnableColor:      true,
		EnableCaller:     false,
		OutputPaths:      []io.Writer{os.Stdout, file},
		ErrorOutputPaths: []io.Writer{file1, file2},
	}
	logger.Init(option)
	logrus.Debug("你好")
	logrus.WithField("name", "小白").Info("不好")
	logrus.WithField("name", "小白").Error("不好")
}
