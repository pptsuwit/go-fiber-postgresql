package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"go-fiber-crud/app/model"
	"go-fiber-crud/app/utils/errs"
	"math/rand"
	"os"
	"path/filepath"

	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
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

func ImageResizing() {}

func RandomImageName(name string) string {
	extension := GetExtension(name)
	randName := base64.StdEncoding.EncodeToString([]byte(name))
	randName += randStringBytesMaskImpr(3)
	return randName + extension
}
func GetExtension(name string) string {
	return filepath.Ext(name)
}
func randStringBytesMaskImpr(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func CheckFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return errors.Is(error, os.ErrNotExist)
}

func CheckFolderExists(filePath string) (string, error) {
	// path := filepath.Join(viper.GetString("app.asset"), filePath)
	path := fmt.Sprintf("./%s/%s", viper.GetString("app.asset"), filePath)

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}
	return path, nil
}

func GetFilePath(path, name string) string {
	return fmt.Sprintf("%s/%s", path, name)
}
func GetHostPath(host, folder, name string) string {
	return fmt.Sprintf("%s/api/%s/%s/%s", host, viper.GetString("app.asset"), folder, name)
}

func DeleteFile(path string) {
	os.Remove(path)
}
