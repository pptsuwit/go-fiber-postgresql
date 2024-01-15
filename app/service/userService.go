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
type UserService interface {
	GetUsers(model.Pagination) (model.UserResponseWithPagination, error)
	GetUser(id int) (*model.UserResponse, error)
	CreateUser(user *model.UserRequest) (*model.UserResponse, error)
	UpdateUser(id int, user *model.UserRequest) (*model.UserResponse, error)
	DeleteUser(id int) error
}

func NewUserService(repository repository.UserRepository) userService {
	return userService{repository: repository}
}

func (s userService) GetUsers(page model.Pagination) (model.UserResponseWithPagination, error) {
	entities, err, count := s.repository.GetAll(page)
	if err != nil {
		logs.Error(err)
		return model.UserResponseWithPagination{}, err
	}

	responseEntity := []model.UserResponse{}
	for _, item := range entities {
		responseEntity = append(responseEntity, model.UserResponse{
			ID:        item.ID,
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Username:  item.Username,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
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
func (s userService) GetUser(id int) (*model.UserResponse, error) {
	user, err := s.repository.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	response := &model.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
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
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	return newUser, nil
}
func (s userService) UpdateUser(id int, user *model.UserRequest) (*model.UserResponse, error) {

	entity, err := s.repository.Update(id, user)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	newUser := &model.UserResponse{
		ID:        entity.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	return newUser, nil
}

func (s userService) DeleteUser(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
