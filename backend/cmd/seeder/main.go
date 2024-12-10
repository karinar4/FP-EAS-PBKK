package main

import (
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/configs"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/database"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/product"
)

func main() {
	if err := configs.Setup("../../.env"); err != nil {
		panic(err)
	}

	// Setup Database
	db, err := database.New()
	if err != nil {
		panic(err)
	}

	product.CategorySeeds(db)
	product.BrandSeeds(db)
	product.ProductSeeds(db)
}
