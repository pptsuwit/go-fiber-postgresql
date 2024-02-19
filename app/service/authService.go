package service

import (
	"go-fiber-crud/app/config/logs"
	"go-fiber-crud/app/model"
	"go-fiber-crud/app/repository"
	"go-fiber-crud/app/utils"
)

type authService struct {
	repository repository.AuthRepository
}
type AuthService interface {
	Login(login *model.Login, asset model.UserAsset) (*model.LoginResponse, error)
	Register(register *model.Register) (*model.UserResponse, error)
}

func NewAuthService(repository repository.AuthRepository) authService {
	return authService{repository: repository}
}

func (s authService) Login(login *model.Login, asset model.UserAsset) (*model.LoginResponse, error) {
	entity, err := s.repository.Login(login)

	if err != nil {
		logs.Error(err)
		return nil, err
	}

	assetFile := ""
	if entity.User.Asset.ID != 0 {
		assetFile = utils.GetHostPath(asset.HostName, asset.FolderName, entity.User.Asset.Name)
	}

	response := &model.LoginResponse{
		Token: entity.Token,
		User: model.UserResponse{
			ID:        entity.User.ID,
			FirstName: entity.User.FirstName,
			LastName:  entity.User.LastName,
			Username:  entity.User.Username,
			CreatedAt: entity.User.CreatedAt,
			UpdatedAt: entity.User.UpdatedAt,
			AssetFile: assetFile,
		},
	}
	return response, nil
}
func (s authService) Register(register *model.Register) (*model.UserResponse, error) {
	response, err := s.repository.Register(register)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return response, nil
}
