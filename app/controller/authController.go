package controller

import (
	"go-fiber-crud/app/model"
	"go-fiber-crud/app/service"
	"go-fiber-crud/app/utils"
	"go-fiber-crud/app/utils/errs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	services service.AuthService
}

func NewAuthController(customerService service.AuthService) authController {
	return authController{
		services: customerService,
	}
}

func (h authController) Login(c *fiber.Ctx) error {
	var request model.Login

	if err := c.BodyParser(&request); err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	validate := validator.New()
	err := validate.Struct(model.Login{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		utils.HandleValidationError(c, err)
		return err
	}
	asset := model.UserAsset{
		HostName:   c.BaseURL(),
		FolderName: "avatars",
	}
	data, err := h.services.Login(&request, asset)
	if err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	utils.ResponseData(c, data)
	return nil
}

func (h authController) Register(c *fiber.Ctx) error {
	var request model.Register

	if err := c.BodyParser(&request); err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	validate := validator.New()
	err := validate.Struct(model.Register{
		Username:  request.Username,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		// Phone:     request.Phone,
	})
	if err != nil {
		utils.HandleValidationError(c, err)
		return err
	}
	data, err := h.services.Register(&request)
	if err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	utils.ResponseData(c, data)
	return nil
}
