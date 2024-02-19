package service

import (
	"go-fiber-crud/app/config/logs"
	"go-fiber-crud/app/model"
	"go-fiber-crud/app/repository"
	"go-fiber-crud/app/utils"
)

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) userService {
	return userService{repository: repository}
}

type UserService interface {
	GetUsers(page model.Pagination, asset model.UserAsset) (model.UserResponseWithPagination, error)
	GetUser(id uint, asset model.UserAsset) (*model.UserResponse, error)

	CreateUser(user *model.UserRequest) (*model.UserResponse, error)
	UpdateUser(id uint, user *model.UpdateUserRequest, asset model.UserAsset) (*model.UserResponse, error)
	DeleteUser(id uint) error
}

func (s userService) GetUsers(page model.Pagination, asset model.UserAsset) (model.UserResponseWithPagination, error) {
	entities, err, count := s.repository.GetAll(page)
	if err != nil {
		logs.Error(err)
		return model.UserResponseWithPagination{}, err
	}
	responseEntity := []model.UserResponse{}
	for _, user := range entities {
		assetFile := ""
		if user.Asset.ID != 0 {
			assetFile = utils.GetHostPath(asset.HostName, asset.FolderName, user.Asset.Name)
		}
		responseEntity = append(responseEntity, model.UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			AssetFile: assetFile,
		})
	}
	response := model.UserResponseWithPagination{
		User: responseEntity,
		Pagination: model.PaginationResponse{
			RecordPerPage: page.PageSize,
			CurrentPage:   page.Page + 1,
			TotalPage:     utils.GetTotalPage(int(count), page.PageSize),
			TotalRecord:   int(count),
		},
	}
	return response, nil
}
func (s userService) GetUser(id uint, asset model.UserAsset) (*model.UserResponse, error) {
	user, err := s.repository.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	assetFile := ""
	if user.Asset.ID != 0 {
		assetFile = utils.GetHostPath(asset.HostName, asset.FolderName, user.Asset.Name)
	}
	response := &model.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		AssetFile: assetFile,
	}
	return response, nil
}
func (s userService) CreateUser(user *model.UserRequest) (*model.UserResponse, error) {

	entity, err := s.repository.Create(user)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	newUser := &model.UserResponse{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Username:  entity.Username,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		AssetFile: user.AssetHostPath,
	}
	return newUser, nil
}
func (s userService) UpdateUser(id uint, user *model.UpdateUserRequest, asset model.UserAsset) (*model.UserResponse, error) {

	entity, err := s.repository.Update(id, user)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	assetFile := ""
	if entity.AssetId != nil {
		assetFile = utils.GetHostPath(asset.HostName, asset.FolderName, entity.Asset.Name)
	}
	if user.AssetHostPath != "" {
		assetFile = user.AssetHostPath
	}
	newUser := &model.UserResponse{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Username:  entity.Username,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		AssetFile: assetFile,
	}
	return newUser, nil
}

func (s userService) DeleteUser(id uint) error {
	err := s.repository.Delete(id)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
