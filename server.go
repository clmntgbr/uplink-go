package main

import (
	"log"
	"uplink-go/config"
	authHandler "uplink-go/handler/auth"
	projectHandler "uplink-go/handler/project"
	userHandler "uplink-go/handler/user"
	"uplink-go/repository"
	"uplink-go/service/auth"
	"uplink-go/service/project"
	"uplink-go/middleware"

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
	projectService := project.New(projectRepo)

	authMiddleware := middleware.NewAuthMiddleware(authService, userRepo)

	userHandler := userHandler.NewUserHandler(userRepo)
	authHandler := authHandler.NewAuthHandler(authService)
	projectHandler := projectHandler.NewProjectHandler(projectService)

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

	api.Use(authMiddleware.Protected())
	api.Use(middleware.InjectUserContext())
	api.Use(middleware.InjectActiveProject(projectRepo))

	api.Get("/user", userHandler.User)

	api.Post("/projects", projectHandler.CreateProject)
	api.Get("/projects", projectHandler.Projects)
	api.Get("/projects/:id", projectHandler.ProjectById)

	log.Fatal(app.Listen(":3000"))
}
