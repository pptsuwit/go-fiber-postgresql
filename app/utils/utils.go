package utils

import (
	"fmt"
	"go-fiber-crud/app/model"
	"go-fiber-crud/app/utils/errs"

	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func ResponseDataList(c *fiber.Ctx, data interface{}, pagination model.PaginationResponse) {
	c.Status(http.StatusOK).JSON(fiber.Map{
		"data":       data,
		"pagination": pagination,
	})
}
func ResponseData(c *fiber.Ctx, data interface{}) {
	c.Status(http.StatusOK).JSON(fiber.Map{
		"data": data,
	})
}

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}
func HandleError(c *fiber.Ctx, err error) {
	switch e := err.(type) {
	case errs.AppError:
		c.Status(e.Code).JSON(fiber.Map{
			"error": e,
		})
		break
	case error:
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
		break
	}
}

func HandleValidationError(c *fiber.Ctx, err error) {
	validErr := []model.ValidatorErr{}
	for _, err := range err.(validator.ValidationErrors) {
		validErr = append(validErr, model.ValidatorErr{
			Field:   err.Field(),
			Value:   err.Value(),
			Tag:     err.Tag(),
			Type:    fmt.Sprintf("%v", err.Type()),
			Param:   err.Param(),
			Message: err.Error(),
		})

	}

	response := model.ErrorWithValidator{
		Code:         http.StatusBadRequest,
		Message:      err.Error(),
		ValidatorErr: validErr,
	}
	c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"error": response,
	})

}

func IsEmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}

func GetTotalPage(count, pageSize int) int {
	totalPage := (count / pageSize)
	if (count % pageSize) > 0 {
		totalPage += 1
	}
	if pageSize == 1 {
		totalPage = count
	}
	return totalPage
}
