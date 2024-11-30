package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/karinar4/FP-EAS-PBKK/internal/configs"
	"github.com/karinar4/FP-EAS-PBKK/internal/database"
	"github.com/karinar4/FP-EAS-PBKK/internal/middleware"
	"github.com/karinar4/FP-EAS-PBKK/internal/modules/auth"
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

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong test",
		})
	})

	var authRepository auth.IAuthRepository = auth.NewAuthRepository(db)
	var authService auth.IAuthUseCase = auth.NewAuthUseCase(authRepository)
	auth.NewAuthHandler(r, authService, "/api/v1/auth")

	if err := r.Run(":" + configs.Config.APP_PORT); err != nil {
		panic(err)
	}
}
