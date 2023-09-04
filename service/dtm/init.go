package main

import (
	"fmt"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/dtm/request"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	port                 int
	dtmHTTPServerAddress string
	healthCheckAddress   string
)

func InitDTM() {
	log.Debugw("InitDTM started")
	dtmServer := viper.GetString("dtm.server")
	log.Debugw("dtm server", "DtmServer", dtmServer)
	startDTMHTTPServer()
}

func startDTMHTTPServer() {
	app := gin.New()
	api := viper.GetString("dtm.http.api")
	port := viper.GetInt("dtm.http.port")
	dtmHTTPServerAddress = fmt.Sprintf("%s:%d%s",
		viper.GetString("dtm.http.host"),
		port,
		api,
	)
	addRoute(app)
	log.Debugw("start dtm http server", "address", dtmHTTPServerAddress)
	err := app.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Errorw("dtm http server failed", "err", err)
		return
	}

}

func addRoute(app *gin.Engine) {
	// 放弃补偿方法路由

	// 添加健康检查路由
	addHealthCheckRoute(app)
	// 添加其他业务路由
	request.AddAbortedCompensateRoute(app)
	request.AddFavoriteActionRoute(app)
	request.AddCommentActionRoute(app)
}

func addHealthCheckRoute(app *gin.Engine) {
	healthCheckAddress = fmt.Sprintf("%s:%d",
		viper.GetString("dtm.health.host"),
		port,
	)

	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
}

func checkServerHealth() bool {
	log.Debugw("checkServerHealth started", "healthCheckAddress", healthCheckAddress)
	client := &http.Client{
		Timeout: time.Second * 2,
	}
	resp, err := client.Get(healthCheckAddress + "/health")
	if err != nil {
		log.Errorw("dtm http server is not ready", "err", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}
