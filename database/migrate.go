package database

import (
	"log"
	"github.com/gtrirf/namoz-vaqtlari/database/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("❌ Faild to run migrations: %v", err)
		return
	}
	log.Println("✅ All tables migrated successfully")
}