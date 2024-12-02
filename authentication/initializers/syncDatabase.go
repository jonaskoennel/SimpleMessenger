package initializers

import (
	"authentication/auth/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
