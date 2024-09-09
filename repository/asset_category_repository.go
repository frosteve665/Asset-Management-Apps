package repository

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/utils/constant"
	"database/sql"
)

type AssetCategoriesRepository interface {
	BaseRepository[model.AssetCategories]
	// BaseRepositoryPaging[model.AssetCategories]
}

type assetcategoriesRepository struct {
	db *sql.DB
}

// Pagination implements AssetCategoriesRepository.
// func (*assetcategoriesRepository) Pagination(requestPaging dto.PaginationQueryParam, query ...string) ([]model.AssetCategories, dto.PaginationResponse, error) {
// 	panic("unimplemented")
// }

func (a *assetcategoriesRepository) Create(payload model.AssetCategories) error {
	_, err := a.db.Exec(constant.ASSET_CATEGORIES_INSERT, payload.Id, payload.Name)
	if err != nil {
		return err
	}
	return nil
}

func (a *assetcategoriesRepository) Get(id string) (model.AssetCategories, error) {
	var assetcategories model.AssetCategories
	row := a.db.QueryRow(constant.ASSET_CATEGORIES_GET, id)
	err := row.Scan(&assetcategories.Id, &assetcategories.Name)
	if err != nil {
		return model.AssetCategories{}, err
	}
	return assetcategories, nil
}

func (a *assetcategoriesRepository) List() ([]model.AssetCategories, error) {
	rows, err := a.db.Query(constant.ASSET_CATEGORIES_LIST)
	if err != nil {
		return nil, err
	}
	var assetcategoriess []model.AssetCategories

	for rows.Next() {
		var assetcategories model.AssetCategories
		err = rows.Scan(&assetcategories.Id, &assetcategories.Name)
		if err != nil {
			panic(err)
		}

		assetcategoriess = append(assetcategoriess, assetcategories)
	}
	return assetcategoriess, nil
}

func (a *assetcategoriesRepository) Update(payload model.AssetCategories) error {
	_, err := a.db.Exec(constant.ASSET_CATEGORIES_UPDATE, payload.Name, payload.Id)
	if err != nil {
		return err
	}
	return nil
}

func (a *assetcategoriesRepository) Delete(id string) error {
	_, err := a.db.Exec(constant.ASSET_CATEGORIES_DELETE, id)
	if err != nil {
		return err
	}
	return nil
}

func NewAssetCategoriesRepository(db *sql.DB) AssetCategoriesRepository {
	return &assetcategoriesRepository{
		db: db,
	}
}
