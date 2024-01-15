package v1

import (
	"go-fiber-crud/app/controller"
	"go-fiber-crud/app/repository"
	"go-fiber-crud/app/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CustomerRouter(router fiber.Router, db *gorm.DB) {

	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)
	router.Get("/customer", customerController.GetCustomers)
	router.Get("/customer/:id", customerController.GetCustomer)
	router.Post("/customer", customerController.CreateCustomer)
	router.Put("/customer/:id", customerController.UpdateCustomer)
	router.Delete("/customer/:id", customerController.DeleteCustomer)

}
