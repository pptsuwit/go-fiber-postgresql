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

const folderName = "avatars"

type userController struct {
	services  service.UserService
	servAsset service.AssetService
}

func NewUserController(userService service.UserService, assetService service.AssetService) userController {
	return userController{
		services:  userService,
		servAsset: assetService,
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
	asset := model.UserAsset{
		HostName:   c.BaseURL(),
		FolderName: folderName,
	}
	data, err := h.services.GetUsers(pagination, asset)
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
	asset := model.UserAsset{
		HostName:   c.BaseURL(),
		FolderName: folderName,
	}
	data, err := h.services.GetUser(uint(id), asset)
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

	// file
	file, fileErr := c.FormFile("file")
	path, err := utils.CheckFolderExists(folderName)
	name := ""
	if fileErr == nil || file != nil {
		if err != nil {
			utils.HandleValidationError(c, err)
			return err
		}

		name = utils.RandomImageName(file.Filename)
		filePath := utils.GetFilePath(path, name)

		if err := c.SaveFile(file, filePath); err != nil {
			utils.HandleValidationError(c, err)
			return err
		}
		newAsset, errAsset := h.servAsset.CreateAsset(&model.AssetRequest{
			Name:     name,
			FileName: file.Filename,
			FileType: utils.GetExtension(file.Filename),
			FilePath: filePath,
			FileSize: strconv.FormatInt(file.Size, 10),
		})
		if errAsset != nil {
			utils.DeleteFile(filePath)
			utils.HandleValidationError(c, fileErr)
			return fileErr
		}
		hostPath := utils.GetHostPath(c.BaseURL(), folderName, newAsset.Name)
		userRequest.AssetId = &newAsset.ID
		userRequest.AssetHostPath = hostPath
	}

	data, err := h.services.CreateUser(&userRequest)
	if err != nil {
		//can't create -> delete file
		if fileErr == nil || file != nil {
			filePath := utils.GetFilePath(path, name)
			utils.DeleteFile(filePath)
			h.servAsset.DeleteAsset(*userRequest.AssetId)
		}

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			// fix duplicate user later
			const duplicate = "Duplicate user accounts were found."
			utils.HandleError(c, errs.NewValidationError(duplicate))
			return err
		}

		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}

	utils.ResponseData(c, data)
	return nil
}

func (h userController) UpdateUser(c *fiber.Ctx) error {

	var userRequest model.UpdateUserRequest
	if err := c.BodyParser(&userRequest); err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}

	validate := validator.New()

	err := validate.Struct(model.UpdateUserRequest{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Username:  userRequest.Username,
		// Password:  userRequest.Password,
	})
	if err != nil {
		utils.HandleError(c, errs.New(err.Error()))
		return err
	}

	// file
	file, fileErr := c.FormFile("file")
	path, err := utils.CheckFolderExists(folderName)
	name := ""
	if fileErr == nil || file != nil {
		if err != nil {
			utils.HandleValidationError(c, err)
			return err
		}

		name = utils.RandomImageName(file.Filename)
		filePath := utils.GetFilePath(path, name)

		if err := c.SaveFile(file, filePath); err != nil {
			utils.HandleValidationError(c, err)
			return err
		}
		newAsset, errAsset := h.servAsset.CreateAsset(&model.AssetRequest{
			Name:     name,
			FileName: file.Filename,
			FileType: utils.GetExtension(file.Filename),
			FilePath: filePath,
			FileSize: strconv.FormatInt(file.Size, 10),
		})
		if errAsset != nil {
			utils.DeleteFile(filePath)
			utils.HandleValidationError(c, fileErr)
			return fileErr
		}
		hostPath := utils.GetHostPath(c.BaseURL(), folderName, newAsset.Name)
		userRequest.AssetId = &newAsset.ID
		userRequest.AssetHostPath = hostPath
	}

	asset := model.UserAsset{
		HostName:   c.BaseURL(),
		FolderName: folderName,
	}
	data, err := h.services.UpdateUser(userRequest.Id, &userRequest, asset)
	if err != nil {
		if fileErr == nil || file != nil {
			filePath := utils.GetFilePath(path, name)
			utils.DeleteFile(filePath)
			h.servAsset.DeleteAsset(*userRequest.AssetId)
		}
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
	err = h.services.DeleteUser(uint(id))
	if err != nil {
		utils.HandleError(c, errs.NewValidationError(err.Error()))
		return err
	}
	utils.ResponseData(c, nil)
	return nil
}
