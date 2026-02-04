package model

import (
	"log"

	"gorm.io/gorm"
)

func registerModels() []any {
	models := []any{
		&UserModel{},
		&UserActivationModel{},
	}
	return models
}

func AutoMigrationModels(db *gorm.DB) {
	for _, model := range registerModels() {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatal(err)
		}
	}
}
