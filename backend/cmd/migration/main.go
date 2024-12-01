package migration

import (
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/auth"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/brand"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/category"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	// Migrate UserModel
	if err := db.AutoMigrate(&auth.UserModel{}); err != nil {
		return err
	}

	// Migrate BrandModel
	if err := db.AutoMigrate(&brand.BrandModel{}); err != nil {
		return err
	}

	// Migrate CategoryModel
	if err := db.AutoMigrate(&category.CategoryModel{}); err != nil {
		return err
	}

	return nil
}
