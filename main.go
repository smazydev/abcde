package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/smazydev/abcde/app/middleware"
	"github.com/smazydev/abcde/app/repositories"
	"github.com/smazydev/abcde/app/routes"
	"github.com/smazydev/abcde/app/services"
	"github.com/smazydev/abcde/app/utils"
)

func main() {
	// Initialize database connection
	db := utils.ConnectDB()
	// Migrate the database schema
	utils.MigrateDB(db)

	// Initialize Fiber app
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	//Initialize repositories.
	userRepo := repositories.NewUserRepository(db)
	businessRepo := repositories.NewBusinessRepository(db)

	//Services
	authService := services.NewAuthService(userRepo)
	authMiddleware := middleware.AuthMiddleware(authService)
	// Routes
	routes.SetupBusinessRoutes(app, db, businessRepo, authService)
	routes.SetupUserRoutes(app, db, userRepo, authService)
	routes.SetupAuthRoutes(app, authService, userRepo)
	// Start the server
	log.Fatal(app.Listen(":3000"))
}
