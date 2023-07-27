package pkg

import (
	"github.com/spf13/pflag"
	"os"
)

var (
	configPath = pflag.StringP("config_path", "c", "/etc/config.ini", "Configuration file.")
	help       = pflag.BoolP("help", "h", false, "Show this help message.")
)

func ParseCommand() string {

	// 读取命令行参数  这个可以写死
	pflag.Parse()
	if *help {
		pflag.Usage()
		os.Exit(0)
	}

	return *configPath
}
