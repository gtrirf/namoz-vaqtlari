package database

import (
	"fmt"
	"log"
	
	"github.com/gtrirf/namoz-vaqtlari/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(dbconfig *config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbconfig.DBHost,
		dbconfig.DBPort,
		dbconfig.DBUsername,
		dbconfig.DBPassword,
		dbconfig.DBName,
	)

	// GORM konfiguratsiyasi
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn), // Faqat errorlarni log qilish
		PrepareStmt: true,           // Prepared statementlar
		SkipDefaultTransaction: true, // Kichik transaksiyalarni optimallashtirish
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Connection pool sozlamalari
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB instance: %v", err)
	}

	// Test query
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	return db
}