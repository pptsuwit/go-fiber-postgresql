package v1

import (
	"go-fiber-crud/app/controller"
	"go-fiber-crud/app/repository"
	"go-fiber-crud/app/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRouter(router fiber.Router, db *gorm.DB) {

	userRepository := repository.NewUserRepositoryDB(db)
	assetRepository := repository.NewAssetRepositoryDB(db)

	userService := service.NewUserService(userRepository)
	assetService := service.NewAssetService(assetRepository)

	userController := controller.NewUserController(userService, assetService)
	router.Get("/user", userController.GetUsers)
	router.Get("/user/:id", userController.GetUser)
	router.Post("/user", userController.CreateUser)
	router.Put("/user", userController.UpdateUser)
	router.Delete("/user/:id", userController.DeleteUser)

}
