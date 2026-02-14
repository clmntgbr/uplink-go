package main

import (
	"log"

    "uplink-go/config"
	"uplink-go/handler/auth"
	"uplink-go/handler/project"
	"uplink-go/handler/user"
	"uplink-go/middleware"
	"uplink-go/repository"
	"uplink-go/service"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	cfg := config.Load()

	db := config.ConnectDatabase(cfg)
	config.AutoMigrate(db)

    userRepo := repository.NewUserRepository(db)
    projectRepo := repository.NewProjectRepository(db)
    authService := auth.NewAuthService(userRepo, projectRepo, cfg)
	authMiddleware := middleware.NewAuthMiddleware(authService, userRepo)
    createProjectService := project.NewCreateProjectService(projectRepo)
    getProjectsService := project.NewGetProjectsService(userRepo)

    userHandler := user.NewUserHandler(userRepo)
	authHandler := auth.NewAuthHandler(authService)
    projectHandler := project.NewProjectHandler(getProjectsService, createProjectService)

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

	api := app.Group("/api")

    api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

    api.Get("/user", authMiddleware.Protected(), userHandler.User)

    api.Post("/projects", authMiddleware.Protected(), projectHandler.CreateProject)
    api.Get("/projects", authMiddleware.Protected(), projectHandler.Projects)

    log.Fatal(app.Listen(":3000"))
}