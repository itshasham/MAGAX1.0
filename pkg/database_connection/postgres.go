package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectPostgres returns a connected *gorm.DB instance or exits on failure
func ConnectPostgres() *gorm.DB {
	dsn := "host=localhost user=postgres password=12345678 dbname=magmox port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}

	// Optional: check ping via raw SQL DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get DB instance: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("❌ PostgreSQL not reachable: %v", err)
	}

	log.Println("✅ PostgreSQL is reachable and connection is alive.")
	return db
}
