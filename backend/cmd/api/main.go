package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/karinar4/FP-EAS-PBKK/backend/cmd/migration"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/configs"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/database"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/middleware"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/auth"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/brand"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/category"
	"github.com/karinar4/FP-EAS-PBKK/backend/internal/modules/product"
)

func main() {
	// Setup configuration
	if err := configs.Setup(".env"); err != nil {
		panic(err)
	}

	// Setup for production
	if configs.Config.ENV_MODE == "production" {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("Production mode")
	}

	// Start the server
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// Setup Database
	db, err := database.New()
	if err != nil {
		panic(err)
	}

	// Run migration for all models
	if err := migration.Migration(db); err != nil {
		panic(err)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong test",
		})
	})

	var authRepository auth.IAuthRepository = auth.NewAuthRepository(db)
	var authService auth.IAuthUseCase = auth.NewAuthUseCase(authRepository)
	auth.NewAuthHandler(r, authService, "/api/v1/auth")

	var brandRepository brand.IBrandRepository = brand.NewBrandRepository(db)
	var brandService brand.IBrandUseCase = brand.NewBrandUseCase(brandRepository)
	brand.NewBrandHandler(r, brandService, "/api/v1/brand")

	var categoryRepository category.ICategoryRepository = category.NewCategoryRepository(db)
	var categoryService category.ICategoryUseCase = category.NewCategoryUseCase(categoryRepository)
	category.NewCategoryHandler(r, categoryService, "/api/v1/category")

	var productRepository product.IProductRepository = product.NewProductRepository(db)
	var productService product.IProductUseCase = product.NewProductUseCase(productRepository)
	product.NewProductHandler(r, productService, "/api/v1/product")

	if err := r.Run(":" + configs.Config.APP_PORT); err != nil {
		panic(err)
	}
}
