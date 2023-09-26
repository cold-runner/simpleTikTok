package dal

import (
	"github.com/cold-runner/simpleTikTok/pkg/db"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/cold-runner/simpleTikTok/service/user/dal/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	log.Init("log-user")
	DB, err = db.NewDB("db-default")
	if err != nil {
		log.Fatalw("init db failed", "err", err)
	}
	log.Infow("Database initialization succeeded.")
	log.Infow("using db:", "dbName", DB.Name())
	DB.AutoMigrate(&model.User{})
	DB.Exec("ALTER TABLE user ADD CONSTRAINT chk_favoriteCount CHECK (favorite_count >= 0);")
	DB.Exec("ALTER TABLE user ADD CONSTRAINT chk_followCount CHECK (follow_count >= 0);")
	DB.Exec("ALTER TABLE user ADD CONSTRAINT chk_followerCount CHECK (follower_count >= 0);")
	DB.Exec("ALTER TABLE user ADD CONSTRAINT chk_totalFavorited CHECK (total_favorited >= 0);")
	DB.Exec("ALTER TABLE user ADD CONSTRAINT chk_workCount CHECK (work_count >= 0);")
}
