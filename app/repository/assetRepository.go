package repository

import (
	"go-fiber-crud/app/model"
	"go-fiber-crud/app/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type assetRepository struct {
	db *gorm.DB
}
type AssetRepository interface {
	GetById(id uint) (*model.Asset, error)
	Create(model.AssetRequest) (*model.Asset, error)
	Update(id uint, asset *model.AssetRequest) (*model.Asset, error)
	Delete(id uint) error
}

func NewAssetRepositoryDB(db *gorm.DB) assetRepository {
	return assetRepository{db: db}
}
func (r assetRepository) GetById(id uint) (*model.Asset, error) {
	entity := model.Asset{}
	tx := r.db.Preload(clause.Associations).First(&entity, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
func (r assetRepository) Create(data model.AssetRequest) (*model.Asset, error) {
	entity := model.Asset{
		Name:     data.Name,
		FileName: data.FileName,
		FileType: data.FileType,
		FilePath: data.FilePath,
		FileSize: data.FileSize,
	}
	tx := r.db.Create(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
func (r assetRepository) Update(id uint, data *model.AssetRequest) (*model.Asset, error) {
	entity := model.Asset{}
	tx := r.db.First(&entity, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	entity.Name = data.Name
	entity.FileName = data.FileName
	entity.FileType = data.FileType
	entity.FilePath = data.FilePath
	entity.FileSize = data.FileSize
	tx = r.db.Save(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
func (r assetRepository) Delete(id uint) (err error) {
	asset := &model.Asset{}
	tx := r.db.Clauses(clause.Returning{}).Unscoped().Delete(asset, id)
	// tx := r.db.Unscoped().Delete(&model.Asset{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	utils.DeleteFile(asset.FilePath)
	return
}
