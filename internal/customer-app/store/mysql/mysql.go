package mysql

import (
	"fmt"
	"github.com/PYxy/go-web/internal/customer-app/store"
	"github.com/PYxy/go-web/pkg/option"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	mysqlFactory store.Factory
	once         sync.Once
)

type mysqlStore struct {
	db *gorm.DB
}

func (m *mysqlStore) CustomerInfoOption() store.CustomerInfoStore {
	//TODO implement me
	return newCustomerInfo(m)
}

func (m *mysqlStore) CustomerGoodOption() store.CustomerGoodStore {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlStore) Close() error {
	db, err := m.db.DB()
	if err != nil {
		//找不到对象就证明不存在就不管
		return nil
	}

	return db.Close()
}

func migrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&store.Customer{}); err != nil {
		return err
	}
	return nil
}

func GetMySQLFactoryOr(opts *option.MysqlOptions) (store.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("cannot get mysql store fatory")
	}

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		dbIns, err = New(opts)

		// uncomment the following line if you need auto migration the given models
		// not suggested in production environment.
		// migrateDatabase(dbIns)
		mysqlFactory = &mysqlStore{dbIns}
		//数据迁移
		if err = migrateDatabase(dbIns); err != nil {
			logrus.Warn("数据迁移异常:", err)
		}

	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}

// New create a new gorm db instance with the given options.
func New(opts *option.MysqlOptions) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s&timeout=4s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.Database,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//需要添加日志  可以加个日志判断
		Logger: opts.DebugOrNot(),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(opts.MaxConnectionLifeTime) * time.Second)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}
