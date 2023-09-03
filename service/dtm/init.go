package favoriteRequest

import (
	"fmt"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var port int
var dtmHTTPServerAddress string
var serverReady = make(chan bool)
var healthCheckAddress string
var errCh = make(chan error)

func InitDTM() {
	log.Debugw("InitDTM started")
	dtmServer := viper.GetString("dtm.server")
	log.Debugw("dtm server", "DtmServer", dtmServer)
	go startDTMHTTPServer()
	log.Debugw("Waiting for server to be ready")
	select {
	case <-serverReady:
		log.Debugw("Server is ready")
	case err := <-errCh:
		log.Fatalw("Server failed", "err", err)
	}
	for {
		time.Sleep(time.Hour)
	}
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
	go func() {
		if err := app.Run(fmt.Sprintf(":%d", port)); err != nil {
			errCh <- err
		}
	}()
	serverReady <- true
}

func addRoute(app *gin.Engine) {
	// 添加健康检查路由
	addHealthCheckRoute(app)
	// 添加其他业务路由
	AddFavoriteActionRoute(app)
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
