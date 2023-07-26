package logger

import (
	"github.com/sirupsen/logrus"

	"testing"
)

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
