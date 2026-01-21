package db

import (
	"log"

	"github.com/Prakharpandey007/paypocket/internal/model"
	"gorm.io/gorm"
)

// Migrate runs GORM AutoMigrate for all DB models.
// Call this after Connect().
func Migrate() {
	if DB == nil {
		log.Fatal("db.Migrate called before db.Connect (DB is nil)")
	}

	if err := migrateWith(DB); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migrations completed successfully")
}

func migrateWith(db *gorm.DB) error {
	// Add models here as the project grows.
	return db.AutoMigrate(
		&model.User{},
		&model.Org{},
		&model.Invite{},
	)
}

