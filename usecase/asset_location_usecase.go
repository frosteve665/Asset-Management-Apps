package usecase

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/repository"
	"fmt"
)

type AssetLocationUsecase interface {
	RegisterNewLocation(bodyRequest model.AssetLocation) error
	SearchLocationById(id string) (model.AssetLocation, error)
	ShowAllLocation() ([]model.AssetLocation, error)
	EditExistedLocation(bodyRequest model.AssetLocation) error
	DeleteSelectedLocation(id string) error
}

type assetLocationUsecase struct {
	repo repository.AssetLocationRepo
}

func (loc *assetLocationUsecase) RegisterNewLocation(bodyRequest model.AssetLocation) error {
	if bodyRequest.Name == "" {
		return fmt.Errorf("location name is required")
	}

	err := loc.repo.Create(bodyRequest)

	if err != nil {
		return fmt.Errorf("failed to add location : %s", err.Error())
	}
	return nil
}

func (loc *assetLocationUsecase) SearchLocationById(id string) (model.AssetLocation, error) {
	return loc.repo.Get(id)
}

func (loc *assetLocationUsecase) ShowAllLocation() ([]model.AssetLocation, error) {
	return loc.repo.List()
}

func (loc *assetLocationUsecase) EditExistedLocation(bodyRequest model.AssetLocation) error {
	_, err := loc.SearchLocationById(bodyRequest.Id)
	if err != nil {
		return fmt.Errorf("can't find location id")
	}

	if bodyRequest.Name == "" {
		return fmt.Errorf("location name is required")
	}

	err = loc.repo.Update(bodyRequest)
	if err != nil {
		return fmt.Errorf("failed to update location : %s", err.Error())
	}

	return nil
}

func (loc *assetLocationUsecase) DeleteSelectedLocation(id string) error {
	_, err := loc.SearchLocationById(id)
	if err != nil {
		return fmt.Errorf("can't find location id")
	}

	err = loc.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete location : %s", err.Error())
	}

	return nil
}

func NewAssetLocationUsecase(repository repository.AssetLocationRepo) AssetLocationUsecase {
	return &assetLocationUsecase{
		repo: repository,
	}
}
