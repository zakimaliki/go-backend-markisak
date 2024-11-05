package helpers

import (
	"markisak/src/configs"
	"markisak/src/models"
)

func Migration() {
	configs.DB.AutoMigrate(&models.User{})
}
