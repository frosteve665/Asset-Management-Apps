package repository

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/utils/constant"
	"database/sql"
)

type AssetLocationRepo interface {
	BaseRepository[model.AssetLocation]
}

type assetLocationRepo struct {
	db *sql.DB
}

func (loc *assetLocationRepo) Create(bodyRequest model.AssetLocation) error {
	_, err := loc.db.Exec(constant.ASSET_LOCATION_INSERT, bodyRequest.Id, bodyRequest.Name)

	if err != nil {
		return err
	}

	return nil
}

func (loc *assetLocationRepo) List() ([]model.AssetLocation, error) {
	rows, err := loc.db.Query(constant.ASSET_LOCATION_LIST)
	if err != nil {
		return nil, err
	}

	var locations []model.AssetLocation
	for rows.Next() {
		var location model.AssetLocation
		err = rows.Scan(&location.Id, &location.Name)

		if err != nil {
			return nil, err
		}

		locations = append(locations, location)
	}

	return locations, nil
}

func (loc *assetLocationRepo) Get(id string) (model.AssetLocation, error) {
	var location model.AssetLocation

	err := loc.db.QueryRow(constant.ASSET_LOCATION_SEARCH, id).Scan(
		&location.Id,
		&location.Name,
	)

	if err != nil {
		return model.AssetLocation{}, nil
	}

	return location, nil
}

func (loc *assetLocationRepo) Update(bodyRequest model.AssetLocation) error {
	_, err := loc.db.Exec(constant.ASSET_LOCATION_UPDATE, bodyRequest.Id, bodyRequest.Name)
	if err != nil {
		return err
	}

	return nil
}

func (loc *assetLocationRepo) Delete(id string) error {
	_, err := loc.db.Exec(constant.ASSET_LOCATION_DELETE, id)
	if err != nil {
		return err
	}

	return nil
}

func NewAssetLocationRepository(db *sql.DB) AssetLocationRepo {
	return &assetLocationRepo{
		db: db,
	}
}
