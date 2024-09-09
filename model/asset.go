package model

import "time"

type Asset struct {
	Id                  string    `json:"id" binding:"required"`
	CategoryId          string    `json:"categoryId" binding:"required"`
	TransactionDetailId any       `json:"transcationDetailId"`
	Name                string    `json:"name" binding:"required,max=100"`
	Description         string    `json:"description"`
	Qty                 int       `json:"qty" binding:"required"`
	ImageUrl            string    `json:"imageUrl"`
	LocationId          string    `json:"locationId" binding:"required"`
	CreatedAt           time.Time `json:"createdAt" binding:"required"`
	AssetDetail         []AssetDetail
}

type AssetDetail struct {
	Id         string `json:"id"`
	AssetId    string `json:"assetId" binding:"required"`
	LocationId string `json:"locationId" binding:"required"`
	Status     int    `json:"status" binding:"required"`
	UpdatedAt  any    `json:"updatedAt"`
	RemovedAt  any    `json:"removedAt"`
}

type AssetPlacement struct {
	Id            string    `json:"id"`
	AsssetId      string    `json:"assetId" binding:"required"`
	CurrentStatus int       `json:"currentStatus" binding:"required"`
	TargetStatus  int       `json:"targetStatus" binding:"required"`
	LocationId    string    `json:"locationId" binding:"required"`
	Qty           int       `json:"qty" binding:"required"`
	UpdatedAt     time.Time `json:"updateAt" binding:"required"`
}

type AssetCategories struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required,max=100"`
}

type AssetLocation struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required,max=100"`
}
