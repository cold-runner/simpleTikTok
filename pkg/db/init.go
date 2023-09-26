package db

import (
	"fmt"
	viperInit "github.com/cold-runner/simpleTikTok/pkg/config"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormopentracing "gorm.io/plugin/opentracing"
	"time"
)

// MySQLConfig 定义 MySQL 数据库的选项.
type MySQLConfig struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
}

// DSN 从 MySQLConfig 返回 DSN.
func (o *MySQLConfig) DSN() string {

	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		o.Username,
		o.Password,
		o.Host,
		o.Database,
		true,
		"Local")
}

// newMySQL 使用给定的选项创建一个新的 gorm 数据库实例.
func newMySQL(config *MySQLConfig) (*gorm.DB, error) {
	logLevel := logger.Silent
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Errorw("load location failed", "err", err)
		return nil, err
	}
	if config.LogLevel != 0 {
		logLevel = logger.LogLevel(config.LogLevel)
	}
	db, err := gorm.Open(mysql.Open(config.DSN()), &gorm.Config{
		Logger:                 logger.Default.LogMode(logLevel),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Now().In(loc)
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns 设置到数据库的最大打开连接数
	sqlDB.SetMaxOpenConns(config.MaxOpenConnections)

	// SetConnMaxLifetime 设置连接可重用的最长时间
	sqlDB.SetConnMaxLifetime(config.MaxConnectionLifeTime)

	// SetMaxIdleConns 设置空闲连接池的最大连接数
	sqlDB.SetMaxIdleConns(config.MaxIdleConnections)

	return db, nil
}

// newMySQLConfig 从 viper 读取配置并返回 MySQLConfig.
// config 是配置名称，例如: "db-user"
func newMySQLConfig(config string) *MySQLConfig {
	viperInit.InitViperConfig()
	dbConfigs := &MySQLConfig{
		Host:                  viper.GetString(config + ".host"),
		Username:              viper.GetString(config + ".username"),
		Password:              viper.GetString(config + ".password"),
		Database:              viper.GetString(config + ".database"),
		MaxIdleConnections:    viper.GetInt(config + ".max-idle-connections"),
		MaxOpenConnections:    viper.GetInt(config + ".max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration(config + ".max-connection-life-time"),
		LogLevel:              viper.GetInt(config + ".log-level"),
	}

	return dbConfigs
}

// NewDB 读取 db 配置，创建 gorm.DB 实例
func NewDB(DBconfig string) (*gorm.DB, error) {
	config := newMySQLConfig(DBconfig)
	DB, err := newMySQL(config)

	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	return DB, err
}
