package controller_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"asetku-bukan-asetmu/delivery/controller"
	"asetku-bukan-asetmu/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockAssetLocationUsecase struct {
	mock.Mock
}

func (loc *mockAssetLocationUsecase) RegisterNewLocation(bodyRequest model.AssetLocation) error {
	args := loc.Called(bodyRequest)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (loc *mockAssetLocationUsecase) SearchLocationById(id string) (model.AssetLocation, error) {
	args := loc.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(model.AssetLocation), args.Error(1)
	}

	return model.AssetLocation{}, nil
}

func (loc *mockAssetLocationUsecase) ShowAllLocation() ([]model.AssetLocation, error) {
	args := loc.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]model.AssetLocation), args.Error(1)
	}

	return nil, nil
}

func (loc *mockAssetLocationUsecase) EditExistedLocation(bodyRequest model.AssetLocation) error {
	args := loc.Called(bodyRequest)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

func (loc *mockAssetLocationUsecase) DeleteSelectedLocation(id string) error {
	args := loc.Called(id)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

type AssetLocationControllerSuite struct {
	suite.Suite
	router          *gin.Engine
	assetLocUsecase *mockAssetLocationUsecase
	controller      *controller.AssetLocationController
}

func (suite *AssetLocationControllerSuite) SetupTest() {
	suite.router = gin.Default()
	suite.assetLocUsecase = new(mockAssetLocationUsecase)
	suite.controller = controller.NewAssetLocationController(suite.router, suite.assetLocUsecase)
}

var dummyBody = []model.AssetLocation{
	{
		Id:   "1",
		Name: "Location 1",
	},
	{
		Id:   "2",
		Name: "Location 2",
	},
}

func (suite *AssetLocationControllerSuite) TestCreateLocation_Success() {
	bodyDummy := dummyBody[0]

	suite.assetLocUsecase.Mock.On("RegisterNewLocation", mock.AnythingOfType("model.AssetLocation")).Return(nil)

	requestBody, _ := json.Marshal(bodyDummy)
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/asset-location/", strings.NewReader(string(requestBody)))

	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	suite.router.ServeHTTP(response, request)

	expectedResponse := json.RawMessage(`{"message":"success add new location","status":201}`)
	expectedResponseBytes, _ := expectedResponse.MarshalJSON()

	assert.Equal(suite.T(), http.StatusCreated, response.Code)
	assert.Equal(suite.T(), expectedResponseBytes, response.Body.Bytes())
}

func (suite *AssetLocationControllerSuite) TestCreateLocation_BadRequests() {
	bodyDummy := dummyBody[0]
	bodyDummy.Id = ""

	suite.assetLocUsecase.Mock.On("RegisterNewLocation", mock.AnythingOfType("model.AssetLocation")).Return(errors.New("Bad Request"))

	requestBody, _ := json.Marshal(bodyDummy)
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/asset-location/", strings.NewReader(string(requestBody)))

	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	suite.router.ServeHTTP(response, request)

	expectedResponse := json.RawMessage(`{"message":"Key: 'AssetLocation.Id' Error:Field validation for 'Id' failed on the 'required' tag"}`)
	expectedResponseBytes, _ := expectedResponse.MarshalJSON()

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), expectedResponseBytes, response.Body.Bytes())
}

func (suite *AssetLocationControllerSuite) TestCreateLocation_ServerError() {
	bodyDummy := dummyBody[0]

	suite.assetLocUsecase.Mock.On("RegisterNewLocation", mock.AnythingOfType("model.AssetLocation")).Return(errors.New("Internal Server Error"))

	requestBody, _ := json.Marshal(bodyDummy)
	request, _ := http.NewRequest(http.MethodPost, "/api/v1/asset-location/", strings.NewReader(string(requestBody)))

	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	suite.router.ServeHTTP(response, request)

	expectedResponse := json.RawMessage(`{"message":"Internal Server Error"}`)
	expectedResponseBytes, _ := expectedResponse.MarshalJSON()

	assert.Equal(suite.T(), http.StatusInternalServerError, response.Code)
	assert.Equal(suite.T(), expectedResponseBytes, response.Body.Bytes())
}

func (suite *AssetLocationControllerSuite) TestListLocation_Success() {
	suite.assetLocUsecase.Mock.On("ShowAllLocation").Return(dummyBody, nil)

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/asset-location/", nil)
	response := httptest.NewRecorder()

	suite.router.ServeHTTP(response, request)

	expectedResult := json.RawMessage(`{"data":[{"id":"1","name":"Location 1"},{"id":"2","name":"Location 2"}],"message":"Success Show All Locations","status":200}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), expectedResultBytes, response.Body.Bytes())
}

func (suite *AssetLocationControllerSuite) TestListLocation_Empty() {
	suite.assetLocUsecase.Mock.On("ShowAllLocation").Return(nil, errors.New("No Content"))

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/asset-location/", nil)
	response := httptest.NewRecorder()

	suite.router.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusNoContent, response.Code)
}

func (suite *AssetLocationControllerSuite) TestListLocation_ServerError() {
	suite.assetLocUsecase.Mock.On("ShowAllLocation").Return(dummyBody, errors.New("Internal Server Error"))

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/asset-location/", nil)
	response := httptest.NewRecorder()

	suite.router.ServeHTTP(response, request)

	expectedResult := json.RawMessage(`{"message":"Internal Server Error"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusInternalServerError, response.Code)
	assert.Equal(suite.T(), expectedResultBytes, response.Body.Bytes())
}

func (suite *AssetLocationControllerSuite) TestGetLocation_Success() {
	bodyReq := dummyBody[0]

	suite.assetLocUsecase.Mock.On("SearchLocationById", bodyReq.Id).Return(bodyReq, nil)

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/asset-location/1", nil)
	response := httptest.NewRecorder()

	suite.router.ServeHTTP(response, request)

	expectedResult := json.RawMessage(`{"data":{"id":"1","name":"Location 1"},"message":"Success Search Location","status":200}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), expectedResultBytes, response.Body.Bytes())
}

func (suite *AssetLocationControllerSuite) TestGetLocation_ServerError() {
	bodyReq := dummyBody[0]

	suite.assetLocUsecase.Mock.On("SearchLocationById", bodyReq.Id).Return(model.AssetLocation{}, errors.New("location not found"))

	request, _ := http.NewRequest(http.MethodGet, "/api/v1/asset-location/1", nil)
	response := httptest.NewRecorder()

	suite.router.ServeHTTP(response, request)

	expectedResult := json.RawMessage(`{"error":"location not found"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusInternalServerError, response.Code)
	assert.Equal(suite.T(), expectedResultBytes, response.Body.Bytes())
}

func (suite *AssetLocationControllerSuite) TestUpdateLocation_Success() {
	bodyReq := dummyBody[0]

	suite.assetLocUsecase.Mock.On("EditExistedLocation", mock.AnythingOfType("model.AssetLocation")).Return(nil)

	requestBody, _ := json.Marshal(bodyReq)
	request, _ := http.NewRequest(http.MethodPut, "/api/v1/asset-location/1", strings.NewReader(string(requestBody)))
	response := httptest.NewRecorder()

	suite.router.ServeHTTP(response, request)

	expectedResult := json.RawMessage(`{"message":"success update existed location","status":200}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), expectedResultBytes, response.Body.Bytes())
}

func (suite *AssetLocationControllerSuite) TestUpdateLocation_BadRequest() {
	bodyReq := dummyBody[0]
	bodyReq.Id = ""

	suite.assetLocUsecase.Mock.On("EditExistedLocation", mock.AnythingOfType("model.AssetLocation")).Return(errors.New("Bad Request"))

	requestBody, _ := json.Marshal(bodyReq)
	request, _ := http.NewRequest(http.MethodPut, "/api/v1/asset-location/1", strings.NewReader(string(requestBody)))
	response := httptest.NewRecorder()

	suite.router.ServeHTTP(response, request)
	fmt.Println(response.Body)

	expectedResult := json.RawMessage(`{"error":"Key: 'AssetLocation.Id' Error:Field validation for 'Id' failed on the 'required' tag"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), expectedResultBytes, response.Body.Bytes())
}

func (suite *AssetLocationControllerSuite) TestUpdateLocation_ServerError() {
	bodyReq := dummyBody[0]

	suite.assetLocUsecase.Mock.On("EditExistedLocation", mock.AnythingOfType("model.AssetLocation")).Return(errors.New("Internal Server Error"))

	requestBody, _ := json.Marshal(bodyReq)
	request, _ := http.NewRequest(http.MethodPut, "/api/v1/asset-location/1", strings.NewReader(string(requestBody)))
	response := httptest.NewRecorder()

	suite.router.ServeHTTP(response, request)
	fmt.Println(response.Body)

	expectedResult := json.RawMessage(`{"error":"Internal Server Error"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusInternalServerError, response.Code)
	assert.Equal(suite.T(), expectedResultBytes, response.Body.Bytes())
}

func (suite *AssetLocationControllerSuite) TestDeleteLocation_Success() {
	bodyReq := dummyBody[0]

	suite.assetLocUsecase.Mock.On("DeleteSelectedLocation", bodyReq.Id).Return(nil)

	request, _ := http.NewRequest(http.MethodDelete, "/api/v1/asset-location/1", nil)
	response := httptest.NewRecorder()

	suite.router.ServeHTTP(response, request)
	fmt.Println(response.Body)

	expectedResult := json.RawMessage(`{"message":"success delete location","status":200}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), expectedResultBytes, response.Body.Bytes())
}

func (suite *AssetLocationControllerSuite) TestDeleteLocation_Failed() {
	bodyReq := dummyBody[0]

	suite.assetLocUsecase.Mock.On("DeleteSelectedLocation", bodyReq.Id).Return(errors.New("failed to delete location"))

	request, _ := http.NewRequest(http.MethodDelete, "/api/v1/asset-location/1", nil)
	response := httptest.NewRecorder()

	suite.router.ServeHTTP(response, request)
	fmt.Println(response.Body)

	expectedResult := json.RawMessage(`{"error":"failed to delete location"}`)
	expectedResultBytes, _ := expectedResult.MarshalJSON()

	assert.Equal(suite.T(), http.StatusInternalServerError, response.Code)
	assert.Equal(suite.T(), expectedResultBytes, response.Body.Bytes())
}

// func (suite *AssetLocationControllerSuite) TestGetLocation_Success() {
// 	fmt.Println(response.Body)
// }

func TestAssetLocationControllerSuite(t *testing.T) {
	suite.Run(t, new(AssetLocationControllerSuite))
}
