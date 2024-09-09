package controller_test

import (
	"asetku-bukan-asetmu/delivery/controller"
	"asetku-bukan-asetmu/model"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockVendorUsecase struct {
	mock.Mock
}

type VendorControllerSuite struct {
	suite.Suite
	router        *gin.Engine
	vendorUsecase *mockVendorUsecase
	controller    *controller.VendorController
}

func (u *mockVendorUsecase) Create(payload model.Vendor) error {
	args := u.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (u *mockVendorUsecase) List() ([]model.Vendor, error) {
	args := u.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]model.Vendor), args.Error(1)
	}

	return nil, nil
}

func (u *mockVendorUsecase) Get(id string) (model.Vendor, error) {
	args := u.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(model.Vendor), args.Error(1)
	}

	return model.Vendor{}, nil
}

func (u *mockVendorUsecase) Update(payload model.Vendor) error {
	args := u.Called(payload)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (u *mockVendorUsecase) Delete(id string) error {
	args := u.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (suite *VendorControllerSuite) SetupTest() {
	suite.router = gin.Default()
	suite.vendorUsecase = new(mockVendorUsecase)
	suite.controller = controller.NewVendorController(suite.router, suite.vendorUsecase)
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

func (suite *VendorControllerSuite) TestCreateSuccess() {
	payload := dummyPayload[0]

	suite.vendorUsecase.Mock.On("Create", payload).Return(nil)

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/vendor", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"status":"success"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusCreated, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestCreateFailBadRequest() {
	payload := dummyPayload[0]
	payload.Name = ""

	suite.vendorUsecase.Mock.On("Create", payload).Return(errors.New("Bad Request"))

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/vendor", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"message":"Key: 'Vendor.Name' Error:Field validation for 'Name' failed on the 'required' tag","status":"fail"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusBadRequest, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestCreateFailServerError() {
	payload := dummyPayload[0]

	suite.vendorUsecase.Mock.On("Create", payload).Return(errors.New("Internal Server Error"))

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/vendor", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"message":"Internal Server Error","status":"fail"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusInternalServerError, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestListSuccess() {
	suite.vendorUsecase.Mock.On("List").Return(dummyPayload, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/vendor", nil)
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"data":[{"id":"1","name":"Vendor 1","address":"Jl. Vendor 1","phone":"08123456789"},{"id":"2","name":"Vendor 2","address":"Jl. Vendor 2","phone":"08123456788"}],"status":"success"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestListFailNotFound() {
	suite.vendorUsecase.Mock.On("List").Return(nil, errors.New("data not found"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/vendor", nil)
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"message":"data not found","status":"fail"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusNotFound, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestListFailServerError() {
	suite.vendorUsecase.Mock.On("List").Return(dummyPayload, errors.New("Internal Server Error"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/vendor", nil)
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"message":"Internal Server Error","status":"fail"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusInternalServerError, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestGetSuccess() {
	payload := dummyPayload[0]

	suite.vendorUsecase.Mock.On("Get", payload.Id).Return(payload, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/vendor/1", nil)
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"data":{"id":"1","name":"Vendor 1","address":"Jl. Vendor 1","phone":"08123456789"},"status":"success"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestGetFail() {
	payload := dummyPayload[0]

	suite.vendorUsecase.Mock.On("Get", payload.Id).Return(model.Vendor{}, errors.New("vendor not found"))

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/vendor/1", nil)
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"message":"vendor not found","status":"fail"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusNotFound, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestUpdateSuccess() {
	payload := dummyPayload[0]
	payload2 := dummyPayload[1]

	suite.vendorUsecase.Mock.On("Get", payload.Id).Return(payload, nil)
	suite.vendorUsecase.Mock.On("Update", payload2).Return(nil)

	reqBody, _ := json.Marshal(dummyPayload[1])
	req, _ := http.NewRequest("PUT", "/api/v1/vendor/1", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"status":"success"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestUpdateFailNotFound() {
	payload := dummyPayload[0]
	payload2 := dummyPayload[1]

	suite.vendorUsecase.Mock.On("Get", payload.Id).Return(model.Vendor{}, errors.New("vendor not found"))
	suite.vendorUsecase.Mock.On("Update", payload2).Return(nil)

	reqBody, _ := json.Marshal(dummyPayload[1])
	req, _ := http.NewRequest("PUT", "/api/v1/vendor/1", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"message":"vendor not found","status":"fail"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusNotFound, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestUpdateFailBadRequest() {
	payload := dummyPayload[0]
	payload2 := dummyPayload[1]
	payload2.Name = ""

	suite.vendorUsecase.Mock.On("Get", payload.Id).Return(payload, nil)
	suite.vendorUsecase.Mock.On("Update", payload2).Return(errors.New("Bad Request"))

	reqBody, _ := json.Marshal(payload2)
	req, _ := http.NewRequest("PUT", "/api/v1/vendor/1", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"message":"Key: 'Vendor.Name' Error:Field validation for 'Name' failed on the 'required' tag","status":"fail"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusBadRequest, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestUpdateFailServerError() {
	payload := dummyPayload[0]
	payload2 := dummyPayload[1]

	suite.vendorUsecase.Mock.On("Get", payload.Id).Return(payload, nil)
	suite.vendorUsecase.Mock.On("Update", payload2).Return(errors.New("Internal Server Error"))

	reqBody, _ := json.Marshal(dummyPayload[1])
	req, _ := http.NewRequest("PUT", "/api/v1/vendor/1", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"message":"Internal Server Error","status":"fail"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusInternalServerError, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestDeleteSuccess() {
	payload := dummyPayload[0]

	suite.vendorUsecase.Mock.On("Get", payload.Id).Return(payload, nil)
	suite.vendorUsecase.Mock.On("Delete", payload.Id).Return(nil)

	req, _ := http.NewRequest("DELETE", "/api/v1/vendor/1", nil)
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"status":"success"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestDeleteFailNotFound() {
	payload := dummyPayload[0]

	suite.vendorUsecase.Mock.On("Get", payload.Id).Return(model.Vendor{}, errors.New("vendor not found"))
	suite.vendorUsecase.Mock.On("Delete", payload.Id).Return(nil)

	req, _ := http.NewRequest("DELETE", "/api/v1/vendor/1", nil)
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"message":"vendor not found","status":"fail"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusNotFound, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func (suite *VendorControllerSuite) TestDeleteFailServerError() {
	payload := dummyPayload[0]

	suite.vendorUsecase.Mock.On("Get", payload.Id).Return(payload, nil)
	suite.vendorUsecase.Mock.On("Delete", payload.Id).Return(errors.New("Internal Server Error"))

	req, _ := http.NewRequest("DELETE", "/api/v1/vendor/1", nil)
	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	expectedResult := json.RawMessage(`{"message":"Internal Server Error","status":"fail"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusInternalServerError, resp.Code)
	assert.Equal(suite.T(), expectedResultBytes, resp.Body.Bytes())
}

func TestVendorControllerSuite(t *testing.T) {
	suite.Run(t, new(VendorControllerSuite))
}
