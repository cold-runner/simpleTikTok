package dal

import (
	"github.com/cold-runner/simpleTikTok/pkg/db"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/social/dal/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	log.Init("log-social")
	DB, err = db.NewDB("db-default")
	if err != nil {
		log.Fatalw("init db failed", "err", err)
	}
	log.Infow("Database initialization succeeded.")
	log.Infow("using db:", "dbName", DB.Name())
	DB.AutoMigrate(&model.UserFavoriteVideo{})
	DB.AutoMigrate(&model.VideoComment{})
}
