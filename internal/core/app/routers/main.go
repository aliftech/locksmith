package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func MainRouter(app *fiber.App, db *gorm.DB) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,X-Requested-With",
		ExposeHeaders:    "Content-Length,Content-Disposition",
		AllowCredentials: false,
		MaxAge:           3600,
	}))
	setupAuthRoutes(app, db)
}
