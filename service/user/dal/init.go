package dal

import (
	"github.com/cold-runner/simpleTikTok/pkg/config"
	"github.com/cold-runner/simpleTikTok/pkg/db"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/user/dal/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	config.InitViperConfig()
	var err error
	log.Init("log-user")
	DB, err = db.NewDB("db-user")
	if err != nil {
		log.Fatalw("db failed", "err", err)
	}
	log.Infow("Database initialization succeeded.")
	log.Infow("using db:", "dbName", DB.Name())
	DB.AutoMigrate(&model.User{})
}
