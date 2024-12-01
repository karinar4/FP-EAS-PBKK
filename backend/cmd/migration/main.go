package migration

import (
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/auth"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/merk"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	// Migrate UserModel
	if err := db.AutoMigrate(&auth.UserModel{}); err != nil {
		return err
	}

	// Migrate MerkModel
	if err := db.AutoMigrate(&merk.MerkModel{}); err != nil {
		return err
	}

	return nil
}
