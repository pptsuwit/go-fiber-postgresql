package repository

import (
	"go-fiber-crud/app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type customerRepository struct {
	db *gorm.DB
}
type CustomerRepository interface {
	GetAll(model.Pagination) ([]model.Customer, error, int64)
	GetById(id int) (*model.Customer, error)
	Create(*model.CustomerRequest) (*model.Customer, error)
	Update(id int, customer *model.CustomerRequest) (*model.Customer, error)
	Delete(id int) error
}

func NewCustomerRepositoryDB(db *gorm.DB) customerRepository {
	return customerRepository{db: db}
}
func (r customerRepository) GetAll(page model.Pagination) ([]model.Customer, error, int64) {
	limit := page.PageSize
	offset := page.Page * limit

	entities := []model.Customer{}
	tx := r.db.Limit(limit).Offset(offset).Preload(clause.Associations).Find(&entities)
	if tx.Error != nil {
		return nil, tx.Error, 0
	}

	// Read
	var countCustomer []model.Customer
	var count int64
	r.db.Model(&countCustomer).Count(&count)
	return entities, nil, count
}
func (r customerRepository) GetById(id int) (*model.Customer, error) {
	entity := model.Customer{}
	tx := r.db.Preload(clause.Associations).First(&entity, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
func (r customerRepository) Create(data *model.CustomerRequest) (*model.Customer, error) {
	entity := model.Customer{
		CustomerId: 001,
		FirstName:  data.FirstName,
		LastName:   data.LastName,
		Address:    data.Address,
		Phone:      data.Phone,
		Email:      data.Email,
		Birthdate:  data.Birthdate,
	}
	tx := r.db.Create(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
func (r customerRepository) Update(id int, data *model.CustomerRequest) (*model.Customer, error) {
	entity := model.Customer{}
	tx := r.db.First(&entity, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	entity.FirstName = data.FirstName
	entity.LastName = data.LastName
	entity.Address = data.Address
	entity.Phone = data.Phone
	entity.Email = data.Email
	entity.Birthdate = data.Birthdate
	tx = r.db.Save(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
func (r customerRepository) Delete(id int) (err error) {
	tx := r.db.Delete(&model.Customer{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return
}
