package dal

import (
	"github.com/cold-runner/simpleTikTok/pkg/db"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/video/dal/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	log.Init("log-user")
	DB, err = db.NewDB("db-default")
	if err != nil {
		log.Fatalw("db failed", "err", err)
	}
	log.Infow("Database initialization succeeded.")
	log.Infow("using db:", "dbName", DB.Name())
	DB.AutoMigrate(&model.Video{})
	DB.Exec("ALTER TABLE video ADD CONSTRAINT chk_favoriteCount_video CHECK (favorite_count >= 0);")
	DB.Exec("ALTER TABLE video ADD CONSTRAINT chk_commentCount_video CHECK (comment_count >= 0);")
}
