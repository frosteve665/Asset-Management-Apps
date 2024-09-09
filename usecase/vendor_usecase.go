package usecase

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/repository"
	"errors"
)

type VendorUsecase interface {
	Create(payload model.Vendor) error
	List() ([]model.Vendor, error)
	Get(id string) (model.Vendor, error)
	Update(payload model.Vendor) error
	Delete(id string) error
}

type vendorUsecase struct {
	repository repository.VendorRepository
}

func (u *vendorUsecase) Create(payload model.Vendor) error {
	return u.repository.Create(payload)
}

func (u *vendorUsecase) List() ([]model.Vendor, error) {
	return u.repository.List()
}

func (u *vendorUsecase) Get(id string) (model.Vendor, error) {
	if id == "" {
		return model.Vendor{}, errors.New("id is required")
	}

	vendor, err := u.repository.Get(id)
	if err != nil {
		return model.Vendor{}, errors.New("vendor not found")
	}

	return vendor, nil
}

func (u *vendorUsecase) Update(payload model.Vendor) error {
	return u.repository.Update(payload)
}

func (u *vendorUsecase) Delete(id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	return u.repository.Delete(id)
}

func NewVendorUsecase(repository repository.VendorRepository) VendorUsecase {
	return &vendorUsecase{repository}
}
