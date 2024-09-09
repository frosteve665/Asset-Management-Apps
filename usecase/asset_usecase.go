package usecase

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/model/dto"
	"asetku-bukan-asetmu/repository"
	"asetku-bukan-asetmu/utils/common"
	"fmt"
)

type AssetUsecase interface {
	CreateNewAsset(bodyRequest model.Asset) error
	ShowAllAsset() ([]dto.AssetDTO, error)
	GetDetailAsset(id string) (dto.AssetDTO, error)
	UpdateAssetLocation(bodyRequest model.AssetPlacement) ([]string, error)
}

type assetUsecase struct {
	repo        repository.AssetRepository
	locUsecase  AssetLocationUsecase
	ctgrUsecase AssetCategoriesUseCase
}

func (a *assetUsecase) CreateNewAsset(bodyRequest model.Asset) error {
	// Check location id
	location, err := a.locUsecase.SearchLocationById(bodyRequest.LocationId)
	if err != nil {
		return fmt.Errorf("location with id %s is not found", bodyRequest.LocationId)
	}

	// Create asset detail
	assetDetails := make([]model.AssetDetail, 0, bodyRequest.Qty)
	for i := 0; i < bodyRequest.Qty; i++ {
		var assetDetail model.AssetDetail
		assetDetail.Id = common.GenerateUUID()
		assetDetail.AssetId = bodyRequest.Id
		assetDetail.LocationId = location.Id
		assetDetail.Status = 1

		assetDetails = append(assetDetails, assetDetail)
	}
	// Fill asset detail
	bodyRequest.AssetDetail = assetDetails

	err = a.repo.Create(bodyRequest)
	if err != nil {
		return fmt.Errorf("failed to register new asset : %v", err)
	}

	return nil
}

func (a *assetUsecase) ShowAllAsset() ([]dto.AssetDTO, error) {
	assets, err := a.repo.List()
	if err != nil {
		return nil, fmt.Errorf("error get list asset : %s", err.Error())
	}

	assetsResponses := make([]dto.AssetDTO, 0, len(assets))
	for _, asset := range assets {
		category, err := a.ctgrUsecase.FindAssetCategoriesById(asset.CategoryId)
		if err != nil {
			return nil, fmt.Errorf("error get category : %s", err.Error())
		}

		var assetRow dto.AssetDTO
		assetRow.Id = asset.Id
		assetRow.TransactionDetailId = asset.TransactionDetailId
		assetRow.Name = asset.Name
		assetRow.Description = asset.Description
		assetRow.ImageUrl = asset.ImageUrl
		assetRow.Qty = asset.Qty
		assetRow.Category = category

		assetsResponses = append(assetsResponses, assetRow)
	}

	return assetsResponses, nil
}

func (a *assetUsecase) GetDetailAsset(id string) (dto.AssetDTO, error) {
	// Asset
	asset, err := a.repo.Detail(id)
	if err != nil {
		return dto.AssetDTO{}, fmt.Errorf("error get asset : %s", err.Error())
	}

	// Category
	category, err := a.ctgrUsecase.FindAssetCategoriesById(asset.CategoryId)
	if err != nil {
		return dto.AssetDTO{}, fmt.Errorf("error get category : %s", err.Error())
	}

	// Asset Detail
	var assetDetailResponse []dto.AssetDetailDTO
	assetDetail, err := a.repo.AssetDetail(asset.Id)
	if err != nil {
		return dto.AssetDTO{}, fmt.Errorf("error get asset detail : %s", err.Error())
	}

	for _, detail := range assetDetail {
		var detailResponse dto.AssetDetailDTO

		location, err := a.locUsecase.SearchLocationById(detail.LocationId)
		if err != nil {
			return dto.AssetDTO{}, fmt.Errorf("error get location : %s", err.Error())
		}

		detailResponse.Id = detail.Id
		detailResponse.Status = detail.Status
		detailResponse.UpdatedAt = detail.UpdatedAt
		detailResponse.Location = location

		assetDetailResponse = append(assetDetailResponse, detailResponse)
	}

	var assetResponse dto.AssetDTO
	assetResponse.Id = asset.Id
	assetResponse.TransactionDetailId = asset.TransactionDetailId
	assetResponse.Name = asset.Name
	assetResponse.Description = asset.Description
	assetResponse.ImageUrl = asset.ImageUrl
	assetResponse.Qty = asset.Qty
	assetResponse.Category = category
	assetResponse.AssetDetail = assetDetailResponse

	return assetResponse, nil
}

func (a *assetUsecase) UpdateAssetLocation(bodyRequest model.AssetPlacement) ([]string, error) {
	currentQty, err := a.repo.CountCurrentQty(bodyRequest.AsssetId, bodyRequest.Qty, bodyRequest.CurrentStatus)
	if err != nil {
		return nil, fmt.Errorf("error check availability asset : %s", err.Error())
	}

	if !(currentQty > bodyRequest.Qty) {
		return nil, nil
	}

	assetId, err := a.repo.GetAvailabilityId(bodyRequest.Qty, bodyRequest.CurrentStatus, bodyRequest.AsssetId)
	if err != nil {
		return nil, nil
	}

	for index, id := range assetId {
		bodyRequest.Id = id
		err := a.repo.UpdateLocation(bodyRequest)
		if err != nil {
			return nil, fmt.Errorf("failed to update asset in looping index(%d) : %s", index, err.Error())
		}
	}

	return assetId, nil
}

func NewAssetUsecase(repo repository.AssetRepository, locationUsecase AssetLocationUsecase, categoryUsecase AssetCategoriesUseCase) AssetUsecase {
	return &assetUsecase{
		repo:        repo,
		locUsecase:  locationUsecase,
		ctgrUsecase: categoryUsecase,
	}
}
