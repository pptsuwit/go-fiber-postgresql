package controller

import (
	"errors"
	"go-fiber-crud/app/model"
	"go-fiber-crud/app/service"
	"go-fiber-crud/app/utils"
	"go-fiber-crud/app/utils/errs"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type userController struct {
	services service.UserService
}

func NewUserController(userService service.UserService) userController {
	return userController{
		services: userService,
	}
}

func (h userController) GetUsers(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 0
	}
	if page > 0 {
		page = page - 1
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		pageSize = 10
	}

	pagination := model.Pagination{
		Page:     page,
		PageSize: pageSize,
	}
	data, err := h.services.GetUsers(pagination)
	if err != nil {
		utils.HandleError(c, errs.NewStatusInternalServerError("Something went wrong. Please try again later"))
		return err
	}

	utils.ResponseDataList(c, data.User, data.Pagination)
	return nil
}

func (h userController) GetUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		utils.HandleError(c, errs.NewValidationError("Invalid Id"))
		return err
	}
	data, err := h.services.GetUser(id)
	if err != nil {
		utils.HandleError(c, errs.NewNotFoundError(err.Error()))
		return err
	}
	utils.ResponseData(c, data)
	return nil
}

func (h userController) CreateUser(c *fiber.Ctx) error {
	var userRequest model.UserRequest

	if err := c.BodyParser(&userRequest); err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	validate := validator.New()

	err := validate.Struct(model.UserRequest{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Username:  userRequest.Username,
		Password:  userRequest.Password,
	})
	if err != nil {
		utils.HandleError(c, errs.New(err.Error()))
		return err
	}
	data, err := h.services.CreateUser(&userRequest)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			utils.HandleError(c, errs.NewValidationError(err.Error()))
			return err
		}

		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	utils.ResponseData(c, data)
	return nil
}

func (h userController) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		utils.HandleError(c, errs.NewValidationError("Invalid Id"))
		return err
	}
	var userRequest model.UserRequest
	if err := c.BodyParser(&userRequest); err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	data, err := h.services.UpdateUser(id, &userRequest)
	if err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	utils.ResponseData(c, data)
	return nil
}
func (h userController) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		utils.HandleError(c, errs.NewValidationError("Invalid Id"))
		return err
	}
	err = h.services.DeleteUser(id)
	if err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	utils.ResponseData(c, nil)
	return nil
}
