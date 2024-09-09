package usecase

import (
	"asetku-bukan-asetmu/repository"
	"fmt"
)

type TestUsecase interface {
	TestUseCase() error
}

type testUsecase struct {
	repo repository.TestRepo
}

func (b *testUsecase) TestUseCase() error {
	err := b.repo.Test()

	if err != nil {
		return fmt.Errorf("error when use usecase : %s", err.Error())
	}
	return nil
}

func NewTestUsecase(repository repository.TestRepo) TestUsecase {
	return &testUsecase{
		repo: repository,
	}
}
