package router

import (
	"go-fiber-crud/app/router/middleware"
	v1 "go-fiber-crud/app/router/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *fiber.App {
	router := fiber.New(fiber.Config{
		// Prefork: true,
		CaseSensitive: true,
		StrictRouting: true,
	})
	router.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	apiV1 := router.Group("/api")
	v1.AuthRouter(apiV1, db)
	apiV1.Get("/abouts", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "abouts",
		})
	})

	//middleware
	apiV1.Use(middleware.Authorize())

	// authentication middleware
	v1.CustomerRouter(apiV1, db)
	v1.UserRouter(apiV1, db)
	apiV1.Get("/abouts-with-middleware", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "abouts-with-middleware",
		})
	})

	apiV1.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
	return router
}
