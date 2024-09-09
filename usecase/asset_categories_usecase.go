package usecase

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/repository"
	"fmt"
)

type AssetCategoriesUseCase interface {
	RegisterNewAssetCategories(payload model.AssetCategories) error
	FindAllAssetCategoriesList() ([]model.AssetCategories, error)
	FindAssetCategoriesById(id string) (model.AssetCategories, error)
	UpdateAssetCategories(payload model.AssetCategories) error
	DeleteAssetCategories(id string) error
	// FindAllAssetCategorie(requesPaging dto.PaginationParam, byNameEmpl string) ([]model.Employee, dto.Paging, error)
}

type assetcategoriesUseCase struct {
	repo repository.AssetCategoriesRepository
}

func (a *assetcategoriesUseCase) RegisterNewAssetCategories(payload model.AssetCategories) error {
	//check attribute nama dan phoneNumber tidak boleh kosong
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	err := a.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create add category : %s", err.Error())
	}
	return nil
}

func (a *assetcategoriesUseCase) FindAllAssetCategoriesList() ([]model.AssetCategories, error) {
	return a.repo.List()
}

func (a *assetcategoriesUseCase) FindAssetCategoriesById(id string) (model.AssetCategories, error) {
	return a.repo.Get(id)
}

func (a *assetcategoriesUseCase) UpdateAssetCategories(payload model.AssetCategories) error {
	_, err := a.FindAssetCategoriesById(payload.Id)
	if err != nil {
		return fmt.Errorf("can't find category id")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	err = a.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update category : %s", err.Error())
	}

	return nil
}

func (a *assetcategoriesUseCase) DeleteAssetCategories(id string) error {
	_, err := a.FindAssetCategoriesById(id)
	if err != nil {
		return err
	}

	err = a.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete category : %s", err.Error())
	}

	return nil
}

// func (a *assetcategoriesUseCase) FindAllAssetCategories(requesPaging dto.PaginationParam, byNameEmpl string) ([]model.AssetCategories, dto.Paging, error){
// 	return a.repo.Paging(requesPaging,byNameEmpl)
// }

func NewAssetCategoriesUseCase(empRepo repository.AssetCategoriesRepository) AssetCategoriesUseCase {
	return &assetcategoriesUseCase{
		repo: empRepo,
	}
}
