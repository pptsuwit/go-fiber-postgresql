package v1

import (
	"go-fiber-crud/app/controller"
	"go-fiber-crud/app/repository"
	"go-fiber-crud/app/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRouter(router fiber.Router, db *gorm.DB) {

	authRepository := repository.NewAuthRepositoryDB(db)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)
	router.Post("/login", authController.Login)
	router.Post("/register", authController.Register)

}
