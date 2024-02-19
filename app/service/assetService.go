package service

import (
	"go-fiber-crud/app/config/logs"
	"go-fiber-crud/app/model"
	"go-fiber-crud/app/repository"
)

type assetService struct {
	repository repository.AssetRepository
}
type AssetService interface {
	GetAssetById(id uint) (*model.AssetResponse, error)
	CreateAsset(asset *model.AssetRequest) (*model.AssetResponse, error)
	UpdateAsset(id uint, asset *model.AssetRequest) (*model.AssetResponse, error)
	DeleteAsset(id uint) error
}

func NewAssetService(repository repository.AssetRepository) assetService {
	return assetService{repository: repository}
}

func (s assetService) GetAssetById(id uint) (*model.AssetResponse, error) {
	asset, err := s.repository.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	response := &model.AssetResponse{
		ID:        asset.ID,
		Name:      asset.Name,
		FileName:  asset.FileName,
		FileType:  asset.FileType,
		FilePath:  asset.FilePath,
		FileSize:  asset.FileSize,
		CreatedAt: asset.CreatedAt,
		UpdatedAt: asset.UpdatedAt,
	}
	return response, nil
}
func (s assetService) CreateAsset(asset *model.AssetRequest) (*model.AssetResponse, error) {

	entity, err := s.repository.Create(*asset)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	newAsset := &model.AssetResponse{
		ID:        entity.ID,
		Name:      asset.Name,
		FileName:  asset.FileName,
		FileType:  asset.FileType,
		FilePath:  asset.FilePath,
		FileSize:  asset.FileSize,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	return newAsset, nil
}
func (s assetService) UpdateAsset(id uint, asset *model.AssetRequest) (*model.AssetResponse, error) {

	entity, err := s.repository.Update(id, asset)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	newAsset := &model.AssetResponse{
		ID:        entity.ID,
		Name:      asset.Name,
		FileName:  asset.FileName,
		FileType:  asset.FileType,
		FilePath:  asset.FilePath,
		FileSize:  asset.FileSize,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	return newAsset, nil
}

func (s assetService) DeleteAsset(id uint) error {
	err := s.repository.Delete(id)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
