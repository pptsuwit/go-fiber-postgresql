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

type customerController struct {
	services service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) customerController {
	return customerController{
		services: customerService,
	}
}

func (h customerController) GetCustomers(c *fiber.Ctx) error {
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
	data, err := h.services.GetCustomers(pagination)
	if err != nil {
		utils.HandleError(c, errs.NewStatusInternalServerError("Something went wrong. Please try again later"))
		return err
	}

	utils.ResponseDataList(c, data.Customer, data.Pagination)
	return nil
}

func (h customerController) GetCustomer(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		utils.HandleError(c, errs.NewValidationError("Invalid Id"))
		return err
	}
	data, err := h.services.GetCustomer(id)
	if err != nil {
		utils.HandleError(c, errs.NewNotFoundError(err.Error()))
		return err
	}
	utils.ResponseData(c, data)
	return nil
}

func (h customerController) CreateCustomer(c *fiber.Ctx) error {
	var customerRequest model.CustomerRequest

	if err := c.BodyParser(&customerRequest); err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	validate := validator.New()

	err := validate.Struct(model.CustomerRequest{
		FirstName: customerRequest.FirstName,
		LastName:  customerRequest.LastName,
		Address:   customerRequest.Address,
		Phone:     customerRequest.Phone,
		Email:     customerRequest.Email,
		Birthdate: customerRequest.Birthdate,
	})
	if err != nil {
		utils.HandleError(c, errs.New(err.Error()))
		return err
	}
	data, err := h.services.CreateCustomer(&customerRequest)
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

func (h customerController) UpdateCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		utils.HandleError(c, errs.NewValidationError("Invalid Id"))
		return err
	}
	var customerRequest model.CustomerRequest
	if err := c.BodyParser(&customerRequest); err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	data, err := h.services.UpdateCustomer(id, &customerRequest)
	if err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	utils.ResponseData(c, data)
	return nil
}
func (h customerController) DeleteCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		utils.HandleError(c, errs.NewValidationError("Invalid Id"))
		return err
	}
	err = h.services.DeleteCustomer(id)
	if err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	utils.ResponseData(c, nil)
	return nil
}
