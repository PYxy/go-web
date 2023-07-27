package pkg

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

var (
	//可以按需添加自定义参数
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

func ParseENV() {
	// 读取环境变量
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)
	key := "keyName"
	if viper.IsSet(key) {
		viper.GetString(key)
	}

}
