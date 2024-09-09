package repository

import (
	"asetku-bukan-asetmu/model"
	"database/sql"
)

type AssetRepository interface {
	Create(bodyRequest model.Asset) error
	List() ([]model.Asset, error)
	Detail(id string) (model.Asset, error)
	AssetDetail(assetId string) ([]model.AssetDetail, error)
	CountCurrentQty(assetId string, qty, currentStatus int) (int, error)
	GetAvailabilityId(limit, status int, assetId string) ([]string, error)
	UpdateLocation(bodyRequest model.AssetPlacement) error
}

type assetRepository struct {
	db *sql.DB
}

func (a *assetRepository) Create(bodyRequest model.Asset) error {
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	// Insert asset
	_, err = tx.Exec("INSERT INTO asset(id,category_id,name,description,image_url,qty,created_at) VALUES($1,$2,$3,$4,$5,$6,$7)", bodyRequest.Id, bodyRequest.CategoryId, bodyRequest.Name, bodyRequest.Description, bodyRequest.ImageUrl, bodyRequest.Qty, bodyRequest.CreatedAt)
	if err != nil {
		return err
	}

	for _, item := range bodyRequest.AssetDetail {
		_, err := tx.Exec("INSERT INTO asset_details(id,asset_id,location_id,status) VALUES($1,$2,$3,$4)", item.Id, item.AssetId, item.LocationId, item.Status)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (a *assetRepository) List() ([]model.Asset, error) {
	rows, err := a.db.Query("SELECT id,category_id,transaction_detail_id,name,description,image_url,qty,created_at FROM asset")
	if err != nil {
		return nil, err
	}

	var assets []model.Asset
	for rows.Next() {
		var asset model.Asset
		err = rows.Scan(&asset.Id, &asset.CategoryId, &asset.TransactionDetailId, &asset.Name, &asset.Description, &asset.ImageUrl, &asset.Qty, &asset.CreatedAt)
		if err != nil {
			return nil, err
		}

		assets = append(assets, asset)
	}

	return assets, nil
}

func (a *assetRepository) Detail(id string) (model.Asset, error) {
	var asset model.Asset
	err := a.db.QueryRow("SELECT id,category_id,transaction_detail_id,name,description,image_url,qty,created_at FROM asset WHERE id=$1", id).Scan(&asset.Id, &asset.CategoryId, &asset.TransactionDetailId, &asset.Name, &asset.Description, &asset.ImageUrl, &asset.Qty, &asset.CreatedAt)
	if err != nil {
		return model.Asset{}, err
	}

	return asset, nil
}

func (a *assetRepository) AssetDetail(assetId string) ([]model.AssetDetail, error) {
	var assetDetails []model.AssetDetail
	rows, err := a.db.Query("SELECT id,location_id,status,updated_at FROM asset_details WHERE asset_id=$1", assetId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var detail model.AssetDetail
		err := rows.Scan(&detail.Id, &detail.LocationId, &detail.Status, &detail.UpdatedAt)
		if err != nil {
			return nil, err
		}

		assetDetails = append(assetDetails, detail)
	}

	return assetDetails, nil
}

func (a *assetRepository) CountCurrentQty(assetId string, qty int, currentStatus int) (int, error) {
	var assetQty int
	err := a.db.QueryRow("SELECT count(*) AS asset_available FROM asset_details WHERE status=$2 AND asset_id=$1", assetId, currentStatus).Scan(&assetQty)
	if err != nil {
		return 0, err
	}

	return assetQty, nil
}

func (a *assetRepository) GetAvailabilityId(limit, status int, assetId string) ([]string, error) {
	var availableAssetId []string
	rows, err := a.db.Query("SELECT id from asset_details WHERE status=$1 AND asset_id=$2 LIMIT $3", status, assetId, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var assetId string
		err := rows.Scan(&assetId)
		if err != nil {
			return nil, err
		}

		availableAssetId = append(availableAssetId, assetId)
	}

	return availableAssetId, nil
}

func (a *assetRepository) UpdateLocation(bodyRequest model.AssetPlacement) error {
	_, err := a.db.Exec("UPDATE asset_details SET location_id=$1, status=$2, updated_at=$3 WHERE id=$4", bodyRequest.LocationId, bodyRequest.TargetStatus, bodyRequest.UpdatedAt, bodyRequest.Id)
	if err != nil {
		return err
	}

	return nil
}

func NewAssetRepository(db *sql.DB) AssetRepository {
	return &assetRepository{
		db: db,
	}
}
