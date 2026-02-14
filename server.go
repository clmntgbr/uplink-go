package main

import (
	"log"

	"uplink-go/config"

  "github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func main() {
		cfg := config.Load()

		db := config.ConnectDatabase(cfg)
		config.AutoMigrate(db)
	
    app := fiber.New(fiber.Config{
        CaseSensitive: true,
        StrictRouting: true,
        ServerHeader:  "Fiber",
        AppName:       "Uplink Go Server v1.0.0",
    })

		app.Use(helmet.New())

		app.Use(logger.New(logger.Config{
				Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}))

		app.Get(healthcheck.LivenessEndpoint, healthcheck.New())
		app.Get(healthcheck.ReadinessEndpoint, healthcheck.New())
		app.Get(healthcheck.StartupEndpoint, healthcheck.New())
		
		// api := app.Group("/api")

		log.Fatal(app.Listen(":3000"))
}