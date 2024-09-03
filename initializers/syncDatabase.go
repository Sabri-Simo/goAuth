package initializers

import "goAuth/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
}
