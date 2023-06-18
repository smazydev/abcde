package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	productRepo := repositories.NewProductRepository(db)

	//initialize services
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	businessService := services.NewBusinessService(businessRepo)
	productService := services.NewProductService(productRepo)

	containerService := services.NewContainer(*userService, authService, *businessService, *productService)

	//Services
	// Routes
	routes.SetupBusinessRoutes(app, db, containerService)
	routes.SetupUserRoutes(app, db, containerService)
	routes.SetupAuthRoutes(app, containerService)
	routes.SetupProductRoutes(app, db, containerService)
	// Start the server
	log.Fatal(app.Listen(":3000"))
}
