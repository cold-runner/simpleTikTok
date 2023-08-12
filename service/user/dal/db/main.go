package main

import (
	"fmt"
	"github.com/cold-runner/simpleTikTok/pkg/init/config"
	"github.com/spf13/viper"
)

func main() {
	config.InitViperConfig()
	fmt.Println("this:", viper.Get("db-user"))
}
