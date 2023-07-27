package main

import (
	"fmt"
	"github.com/PYxy/go-web/pkg/logger"
	_ "github.com/PYxy/go-web/pkg/logger"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func main() {
	//默认日志配置
	//LogByDefault()

	//自定义日志输出
	//LogByOption()

	//日志切割
	LogByRotate()
}

func LogByDefault() {
	logrus.Debug("你好")
	logrus.Info("nihao2")
	logrus.WithField("name", "ljy").Info("lalala")
	logrus.WithField("name", "ljy2").Error("lalala")
}

func LogByOption() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	fmt.Println("文件打开异常:", err)
	defer file.Close()
	file1, err := os.OpenFile("err.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	fmt.Println("文件打开异常:", err)
	defer file1.Close()
	option := &logger.Options{
		Level:            logrus.DebugLevel,
		Format:           "console", //json console
		EnableColor:      true,
		EnableCaller:     true,
		OutputPaths:      []io.Writer{os.Stdout, file},
		ErrorOutputPaths: []io.Writer{file1},
	}
	logger.Init(option)
	logrus.Debug("你好")
	logrus.Info("nihao2")
	logrus.WithField("name", "ljy").Info("lalala")
	logrus.WithField("name", "ljy2").Error("lalala")
}

func LogByRotate() {
	path := "log_rotate.log"
	writer, _ := rotateLogs.New(
		path+".%Y%m%d%H%M",
		rotateLogs.WithClock(rotateLogs.Local),
		rotateLogs.WithLinkName(path),
		rotateLogs.WithMaxAge(time.Hour*1),
		rotateLogs.WithRotationTime(time.Minute*1),
		//rotateLogs.WithRotationCount(6),
	)
	option := &logger.Options{
		Level:        logrus.DebugLevel,
		Format:       "console", //json console
		EnableColor:  true,
		EnableCaller: true,
		OutputPaths:  []io.Writer{os.Stdout, writer},
		//ErrorOutputPaths: []io.Writer{file1},
	}

	logger.Init(option)
	logrus.Debug("你好")
	logrus.Info("nihao2")
	logrus.WithField("name", "ljy").Info("lalala")
	logrus.WithField("name", "ljy2").Error("lalala")

	time.Sleep(time.Minute * 2)
	logrus.Debug("你好")
	logrus.Info("nihao2")
	logrus.WithField("name", "ljy").Info("lalala")
}
