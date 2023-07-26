package config

import (
	"github.com/PYxy/go-web/pkg/config"
	"github.com/PYxy/go-web/pkg/option"
)

type ConfigOption struct {
	MysqlOption *option.MysqlOptions `ini:"mysql"`
	//RedisOption option.RedisOptions `ini:"redis"`
}

func ParseAppInI(filePath string) (*ConfigOption, error) {
	var configOption *ConfigOption
	iniFile, err := config.ReadInI(filePath)
	if err != nil {
		return nil, err
	}
	configOption = &ConfigOption{}
	err = iniFile.MapTo(&configOption)

	return configOption, err
}
