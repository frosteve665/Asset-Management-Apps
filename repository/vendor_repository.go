package repository

import (
	"asetku-bukan-asetmu/model"
	"database/sql"
)

type VendorRepository interface {
	BaseRepository[model.Vendor]
}

type vendorRepository struct {
	db *sql.DB
}

func (r *vendorRepository) Create(payload model.Vendor) error {
	_, err := r.db.Exec("INSERT INTO vendors (id, name, address, phone) VALUES ($1, $2, $3, $4)", payload.Id, payload.Name, payload.Address, payload.Phone)
	return err
}

func (r *vendorRepository) List() ([]model.Vendor, error) {
	rows, err := r.db.Query("SELECT id, name, address, phone FROM vendors")
	if err != nil {
		return nil, err
	}

	var vendors []model.Vendor
	for rows.Next() {
		var vendor model.Vendor
		if err := rows.Scan(&vendor.Id, &vendor.Name, &vendor.Address, &vendor.Phone); err != nil {
			return nil, err
		}
		vendors = append(vendors, vendor)
	}

	return vendors, nil
}

func (r *vendorRepository) Get(id string) (model.Vendor, error) {
	var vendor model.Vendor
	err := r.db.QueryRow("SELECT id, name, address, phone FROM vendors WHERE id = $1", id).Scan(&vendor.Id, &vendor.Name, &vendor.Address, &vendor.Phone)
	return vendor, err
}

func (r *vendorRepository) Update(payload model.Vendor) error {
	_, err := r.db.Exec("UPDATE vendors SET name = $1, address = $2, phone = $3 WHERE id = $4", payload.Name, payload.Address, payload.Phone, payload.Id)
	return err
}

func (r *vendorRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM vendors WHERE id = $1", id)
	return err
}

func NewVendorRepository(db *sql.DB) VendorRepository {
	return &vendorRepository{db}
}
