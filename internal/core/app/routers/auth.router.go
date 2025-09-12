package routers

import (
	"github.com/aliftech/locksmith/internal/core/app/controllers"
	"github.com/aliftech/locksmith/internal/core/app/repositories/mysql"
	"github.com/aliftech/locksmith/internal/core/app/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func setupAuthRoutes(app *fiber.App, db *gorm.DB) {
	userAuthRepo := mysql.NewAuthRepository(db)
	userRepo := mysql.NewUserRepository(db)
	authService := services.NewAuthService(userRepo, userAuthRepo)
	authCtrl := controllers.NewAuthController(authService)

	app.Post("/api/v1/auth/signup", authCtrl.HandleSignup)
}
