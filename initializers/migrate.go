package initializers

import "Go-Gin/models"

func Migrate() {
	Db.AutoMigrate(&models.User{}, &models.Post{},)
}
