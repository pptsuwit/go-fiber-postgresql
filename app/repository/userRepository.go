package repository

import (
	"go-fiber-crud/app/model"
	"go-fiber-crud/app/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}
type UserRepository interface {
	GetAll(model.Pagination) ([]model.User, error, int64)
	GetById(id uint) (*model.User, error)
	Create(*model.UserRequest) (*model.User, error)
	Update(id uint, user *model.UpdateUserRequest) (*model.User, error)
	Delete(id uint) error
	DeleteAsset(id uint) (err error)
}

func NewUserRepositoryDB(db *gorm.DB) userRepository {
	return userRepository{db: db}
}
func (r userRepository) GetAll(page model.Pagination) ([]model.User, error, int64) {
	limit := page.PageSize
	offset := page.Page * limit

	entities := []model.User{}
	tx := r.db.Limit(limit).Offset(offset).Preload(clause.Associations).Order("id desc").Find(&entities)
	if tx.Error != nil {
		return nil, tx.Error, 0
	}

	// Read
	var countUser []model.User
	var count int64
	r.db.Model(&countUser).Count(&count)
	return entities, nil, count
}
func (r userRepository) GetById(id uint) (*model.User, error) {
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
		AssetId:   data.AssetId,
	}
	tx := r.db.Create(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
func (r userRepository) Update(id uint, data *model.UpdateUserRequest) (*model.User, error) {
	entity := model.User{}
	tx := r.db.First(&entity, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, err
	// }

	entity.FirstName = data.FirstName
	entity.LastName = data.LastName
	entity.Username = data.Username
	// entity.Password = string(hashedPassword)

	if data.AssetId != nil {
		r.DeleteAsset(entity.ID)
		entity.AssetId = data.AssetId
	}
	tx = r.db.Save(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}

	r.db.Preload(clause.Associations).First(&entity, id)
	return &entity, nil
}
func (r userRepository) Delete(id uint) (err error) {
	r.DeleteAsset(id)
	tx := r.db.Preload(clause.Associations).Delete(&model.User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return
}

func (r userRepository) DeleteAsset(id uint) (err error) {

	entity := model.User{}
	tx := r.db.Clauses(clause.Returning{}).Preload(clause.Associations).First(&entity, id)
	if tx.Error != nil {
		return tx.Error
	}

	// r.db.Unscoped().Delete(&model.Asset{}, entity.AssetId)
	asset := &model.Asset{}
	ax := r.db.Clauses(clause.Returning{}).Unscoped().Delete(asset, entity.AssetId)
	if ax.Error != nil {
		return tx.Error
	}
	utils.DeleteFile(asset.FilePath)
	return
}
