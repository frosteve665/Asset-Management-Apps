package dto

import (
	"asetku-bukan-asetmu/model"
)

type AssetDTO struct {
	Id                  string                `json:"id"`
	TransactionDetailId any                   `json:"transactionDetailId"`
	Name                string                `json:"name"`
	Description         string                `json:"description"`
	ImageUrl            string                `json:"imageUrl"`
	Qty                 int                   `json:"qty"`
	Category            model.AssetCategories `json:"category"`
	AssetDetail         []AssetDetailDTO      `json:"assetDetail"`
}

type AssetDetailDTO struct {
	Id        string              `json:"id"`
	Status    int                 `json:"status"`
	Location  model.AssetLocation `json:"location"`
	UpdatedAt any                 `json:"updatedAt"`
}
