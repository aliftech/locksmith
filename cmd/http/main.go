package main

import (
	"os"

	"github.com/aliftech/locksmith/internal/core/app/routers"
	"github.com/aliftech/locksmith/internal/core/config"
	"github.com/gofiber/fiber/v2"
)

func init() {
	config.EnvSetup()
}

func main() {
	app := fiber.New()
	db := config.ConnectDB()

	routers.MainRouter(app, db)
	app.Listen(os.Getenv("TCP"))
}
