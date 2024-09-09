package manager

import (
	"asetku-bukan-asetmu/usecase"
)

type UseCaseManager interface {
	TestUsecase() usecase.TestUsecase
	EmployeeUseCase() usecase.EmployeeUseCase
	AssetUsecase() usecase.AssetUsecase
	AssetLocationUsecase() usecase.AssetLocationUsecase
	AssetCategoriesUseCase() usecase.AssetCategoriesUseCase
	VendorUseCase() usecase.VendorUsecase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) TestUsecase() usecase.TestUsecase {
	return usecase.NewTestUsecase(u.repoManager.TestRepoManager())
}

func (u *useCaseManager) EmployeeUseCase() usecase.EmployeeUseCase {
	return usecase.NewEmployeeUseCase(u.repoManager.EmployeeRepo())
}

func (u *useCaseManager) AssetUsecase() usecase.AssetUsecase {
	return usecase.NewAssetUsecase(u.repoManager.AssetRepo(), u.AssetLocationUsecase(), u.AssetCategoriesUseCase())
}

func (u *useCaseManager) AssetCategoriesUseCase() usecase.AssetCategoriesUseCase {
	return usecase.NewAssetCategoriesUseCase(u.repoManager.AssetCategoriesRepo())
}

func (u *useCaseManager) AssetLocationUsecase() usecase.AssetLocationUsecase {
	return usecase.NewAssetLocationUsecase(u.repoManager.AssetLocationRepo())
}

func (u *useCaseManager) VendorUseCase() usecase.VendorUsecase {
	return usecase.NewVendorUsecase(u.repoManager.VendorRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repo,
	}
}
