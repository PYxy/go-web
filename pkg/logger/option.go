package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

const (
	consoleFormat = "console"
	jsonFormat    = "json"
)

// Options contains configuration items related to log.
type Options struct {
	Level            logrus.Level `json:"level" mapstructure:"level"`
	Format           string       `json:"format" mapstructure:"format"`
	EnableColor      bool         `json:"enable-color" mapstructure:"enable-color"`
	EnableCaller     bool         `json:"enable-caller" mapstructure:"enable-caller"`
	OutputPaths      []io.Writer  `json:"output-paths" mapstructure:"output-paths"`
	ErrorOutputPaths []io.Writer  `json:"error-output-paths" mapstructure:"error-output-paths"`
}

// NewOptions creates a Options object with default parameters.
func NewOptions() *Options {
	return &Options{
		Level:            logrus.InfoLevel,
		Format:           consoleFormat,
		EnableColor:      true,
		EnableCaller:     false,
		OutputPaths:      []io.Writer{os.Stdout},
		ErrorOutputPaths: []io.Writer{os.Stderr},
	}
}
