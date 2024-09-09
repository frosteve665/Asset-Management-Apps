package usecase_test

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockVendorRepository struct {
	mock.Mock
}

type VendorUsecaseTestSuite struct {
	suite.Suite
	mockRepo *mockVendorRepository
	usecase  usecase.VendorUsecase
}

func (r *mockVendorRepository) Create(payload model.Vendor) error {
	args := r.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (r *mockVendorRepository) List() ([]model.Vendor, error) {
	args := r.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]model.Vendor), args.Error(1)
	}

	return nil, nil
}

func (r *mockVendorRepository) Get(id string) (model.Vendor, error) {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(model.Vendor), args.Error(1)
	}

	return model.Vendor{}, nil
}

func (r *mockVendorRepository) Update(payload model.Vendor) error {
	args := r.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (r *mockVendorRepository) Delete(id string) error {
	args := r.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (suite *VendorUsecaseTestSuite) SetupTest() {
	suite.mockRepo = new(mockVendorRepository)
	suite.usecase = usecase.NewVendorUsecase(suite.mockRepo)
}

var dummyPayload = []model.Vendor{
	{
		Id:      "1",
		Name:    "Vendor 1",
		Address: "Jl. Vendor 1",
		Phone:   "08123456789",
	},
	{
		Id:      "2",
		Name:    "Vendor 2",
		Address: "Jl. Vendor 2",
		Phone:   "08123456788",
	},
}

func (suite *VendorUsecaseTestSuite) TestCreateSuccess() {
	payload := dummyPayload[0]
	suite.mockRepo.Mock.On("Create", payload).Return(nil)

	err := suite.usecase.Create(payload)
	assert.NoError(suite.T(), err)
}

func (suite *VendorUsecaseTestSuite) TestCreateFail() {
	payload := dummyPayload[0]
	suite.mockRepo.Mock.On("Create", payload).Return(assert.AnError)

	err := suite.usecase.Create(payload)
	assert.Error(suite.T(), err)
}

func (suite *VendorUsecaseTestSuite) TestListSuccess() {
	suite.mockRepo.Mock.On("List").Return(dummyPayload, nil)

	result, err := suite.usecase.List()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyPayload, result)
}

func (suite *VendorUsecaseTestSuite) TestListFail() {
	suite.mockRepo.Mock.On("List").Return([]model.Vendor{}, assert.AnError)

	result, err := suite.usecase.List()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), []model.Vendor{}, result)
}

func (suite *VendorUsecaseTestSuite) TestGetSuccess() {
	expectedResult := dummyPayload[0]

	suite.mockRepo.Mock.On("Get", expectedResult.Id).Return(expectedResult, nil)

	result, err := suite.usecase.Get(expectedResult.Id)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedResult, result)
}

func (suite *VendorUsecaseTestSuite) TestGetFail() {
	expectedResult := dummyPayload[0]

	suite.mockRepo.Mock.On("Get", expectedResult.Id).Return(model.Vendor{}, assert.AnError)

	result, err := suite.usecase.Get(expectedResult.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Vendor{}, result)
	assert.Equal(suite.T(), "vendor not found", err.Error())

	result, err = suite.usecase.Get("")
	assert.Equal(suite.T(), "id is required", err.Error())
}

func (suite *VendorUsecaseTestSuite) TestUpdateSuccess() {
	payload := dummyPayload[0]

	suite.mockRepo.Mock.On("Update", payload).Return(nil)

	err := suite.usecase.Update(payload)
	assert.NoError(suite.T(), err)
}

func (suite *VendorUsecaseTestSuite) TestUpdateFail() {
	payload := dummyPayload[0]

	suite.mockRepo.Mock.On("Update", payload).Return(assert.AnError)

	err := suite.usecase.Update(payload)
	assert.Error(suite.T(), err)
}

func (suite *VendorUsecaseTestSuite) TestDeleteSuccess() {
	suite.mockRepo.On("Delete", dummyPayload[0].Id).Return(nil)

	err := suite.usecase.Delete(dummyPayload[0].Id)
	assert.NoError(suite.T(), err)
}

func (suite *VendorUsecaseTestSuite) TestDeleteFail() {
	suite.mockRepo.On("Delete", dummyPayload[0].Id).Return(assert.AnError)

	err := suite.usecase.Delete(dummyPayload[0].Id)
	assert.Error(suite.T(), err)

	err = suite.usecase.Delete("")
	assert.Equal(suite.T(), "id is required", err.Error())
}

func TestVendorUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(VendorUsecaseTestSuite))
}
