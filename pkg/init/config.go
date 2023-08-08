package main

import (
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type LogConfig struct {
	DisableCaller     bool     `yaml:"disable-caller"`
	DisableStacktrace bool     `yaml:"disable-stacktrace"`
	Level             string   `yaml:"level"`
	Format            string   `yaml:"format"`
	OutputPaths       []string `yaml:"output-paths"`
}

// 初始化viper配置
func initViperConfig() {
	// 获取当前文件的路径
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatalw("Failed to get working directory:", "err", err)
	}
	// 设置配置文件的路径
	configPath := filepath.Join(projectRoot, "config", "config.yaml")
	log.Infow("Reading configuration file at:", "path", configPath)
	// 设置viper相关配置
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// 读取配置文件
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalw("Failed to read the configuration file: %v", "err", err)
	}
}

// 传入的是一个字符串切片用于指定配置名，返回一个log.Options结构体，用于初始化log
func logOptions(logOpt string) *log.Options {
	if logOpt != "" {
		return &log.Options{
			DisableCaller: viper.GetBool(logOpt +
				".disable-caller"),
			DisableStacktrace: viper.GetBool(logOpt + ".disable-stacktrace"),
			Level:             viper.GetString(logOpt + ".level"),
			Format:            viper.GetString(logOpt + ".format"),
			OutputPaths:       viper.GetStringSlice(logOpt + ".output-paths"),
		}
	}
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}

}

func main() {
	initViperConfig()
	opt := logOptions("log-user")
	log.Init(opt)
	log.Infow("hello world", "Info", "Test")
	defer log.Sync()
}
