package repository

import (
	"go-fiber-crud/app/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}
type UserRepository interface {
	GetAll(model.Pagination) ([]model.User, error, int64)
	GetById(id int) (*model.User, error)
	Create(*model.UserRequest) (*model.User, error)
	Update(id int, user *model.UserRequest) (*model.User, error)
	Delete(id int) error
}

func NewUserRepositoryDB(db *gorm.DB) userRepository {
	return userRepository{db: db}
}
func (r userRepository) GetAll(page model.Pagination) ([]model.User, error, int64) {
	limit := page.PageSize
	offset := page.Page * limit

	entities := []model.User{}
	tx := r.db.Limit(limit).Offset(offset).Preload(clause.Associations).Find(&entities)
	if tx.Error != nil {
		return nil, tx.Error, 0
	}

	// Read
	var countUser []model.User
	var count int64
	r.db.Model(&countUser).Count(&count)
	return entities, nil, count
}
func (r userRepository) GetById(id int) (*model.User, error) {
	entity := model.User{}
	tx := r.db.Preload(clause.Associations).First(&entity, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
func (r userRepository) Create(data *model.UserRequest) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	entity := model.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Username:  data.Username,
		Password:  string(hashedPassword),
	}
	tx := r.db.Create(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
func (r userRepository) Update(id int, data *model.UserRequest) (*model.User, error) {
	entity := model.User{}
	tx := r.db.First(&entity, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	entity.FirstName = data.FirstName
	entity.LastName = data.LastName
	entity.Username = data.Username
	entity.Password = string(hashedPassword)
	tx = r.db.Save(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
func (r userRepository) Delete(id int) (err error) {
	tx := r.db.Delete(&model.User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return
}
