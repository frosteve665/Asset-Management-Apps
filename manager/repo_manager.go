package manager

import "asetku-bukan-asetmu/repository"

type RepoManager interface {
	TestRepoManager() repository.TestRepo
	EmployeeRepo() repository.EmployeeRepository
	AssetRepo() repository.AssetRepository
	AssetCategoriesRepo() repository.AssetCategoriesRepository
	AssetLocationRepo() repository.AssetLocationRepo
	VendorRepo() repository.VendorRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) TestRepoManager() repository.TestRepo {
	return repository.NewTestRepository(r.infra.Connection())
}

func (r *repoManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmployeeRepository(r.infra.Connection())
}

func (r *repoManager) AssetRepo() repository.AssetRepository {
	return repository.NewAssetRepository(r.infra.Connection())
}

func (r *repoManager) AssetCategoriesRepo() repository.AssetCategoriesRepository {
	return repository.NewAssetCategoriesRepository(r.infra.Connection())
}

func (r *repoManager) AssetLocationRepo() repository.AssetLocationRepo {
	return repository.NewAssetLocationRepository(r.infra.Connection())
}

func (r *repoManager) VendorRepo() repository.VendorRepository {
	return repository.NewVendorRepository(r.infra.Connection())
}

func NewRepoManager(infraParam InfraManager) RepoManager {
	return &repoManager{
		infra: infraParam,
	}
}
