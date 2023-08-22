package config

import (
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

// 初始化viper配置
func InitViperConfig() {
	// 获取当前文件的路径
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatalw("Failed to get working directory:", "err", err)
	}

	// 查找 "simpleTikTok" 目录，并截取该路径
	index := strings.Index(projectRoot, "simpleTikTok")
	if index == -1 {
		log.Fatalw("Cannot find 'simpleTikTok' directory in path")
	}

	// 构造配置文件的路径
	configPath := filepath.Join(projectRoot[:index+len("simpleTikTok")], "config", "config.yaml")
	log.Infow("Reading configuration file at:", "path", configPath)
	// 设置viper相关配置
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// 读取配置文件
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalw("Failed to read the configuration file: %v", "err", err)
	}
}

//func main() {
//	initViperConfig()
//	opt := logOptions("log-user")
//	log.Init(opt)
//	log.Infow("hello world", "Info", "Test")
//	defer log.Sync()
//}
