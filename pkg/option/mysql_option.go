package option

import (
	"gorm.io/gorm/logger"
	slog "log"
	"os"
	"time"
)

// MysqlOptions 连接信息 以及mysql 连接对象创建
type MysqlOptions struct {
	Host                  string `ini:"host"`
	Port                  int    `ini:"port"`
	Username              string `ini:"username"`
	Password              string `ini:"password"`
	Database              string `ini:"database"`
	MaxIdleConnections    int    `ini:"maxIdleConnections"`
	MaxOpenConnections    int    `ini:"maxOpenConnections"`
	MaxConnectionLifeTime int    `ini:"maxConnectionLifeTime"`
	// DebugOr 是否打印日志
	DebugOr bool `ini:"debug"`
	Log     logger.Interface
}

type mysqlOption func(options *MysqlOptions)

func WithLog(log logger.Interface) mysqlOption {
	return func(options *MysqlOptions) {
		options.Log = log
	}
}

// DebugOrNot  Based on the DebugOr field  whether to output logs
func (m *MysqlOptions) DebugOrNot() logger.Interface {
	if m.DebugOr {
		//如果用户想打印debug 日志,但是没有设置输出就默认
		if m.Log == nil {
			return logger.New(
				slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
				logger.Config{
					SlowThreshold: time.Second, // 慢查询 SQL 阈值
					Colorful:      false,       // 禁用彩色打印
					//IgnoreRecordNotFoundError: false,
					LogLevel: logger.Info, // Log lever
				},
			)

		}
		return m.Log
	}
	return nil
}
