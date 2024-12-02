package initializers

import (
	"SimpleMessenger/m/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
