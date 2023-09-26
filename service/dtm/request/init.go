package request

import (
	"fmt"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	api       string
	host      string
	dtmServer string
)

func InitRequestVariables() {
	host = fmt.Sprintf(
		"%s:%d%s",
		viper.GetString("dtm.http.host"),
		viper.GetInt("dtm.http.port"),
		viper.GetString("dtm.http.api"),
	)

	dtmServer = viper.GetString("dtm.server")
	api = viper.GetString("dtm.http.api")

}

func AddAbortedCompensateRoute(app *gin.Engine) {
	app.POST(api+abortedCompensate, func(c *gin.Context) {
		log.Debugw("Aborted compensate")
		c.JSON(200, "Aborted compensate")
	})
}
