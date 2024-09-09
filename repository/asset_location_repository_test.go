package repository_test

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/repository"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AssetLocationRepositorySuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo repository.AssetLocationRepo
}

func (loc *AssetLocationRepositorySuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(loc.T(), err)

	loc.db = db
	loc.mock = mock
	loc.repo = repository.NewAssetLocationRepository(db)
}

func (loc *AssetLocationRepositorySuite) TearDownTest() {
	loc.db.Close()
}

func (loc *AssetLocationRepositorySuite) TestCreate_Success() {
	bodyRequest := model.AssetLocation{
		Id:   "1",
		Name: "Location1",
	}

	loc.mock.ExpectExec("INSERT INTO asset_location").WithArgs(bodyRequest.Id, bodyRequest.Name).WillReturnResult(sqlmock.NewResult(1, 1))

	err := loc.repo.Create(bodyRequest)
	assert.NoError(loc.T(), err)
}

func (loc *AssetLocationRepositorySuite) TestCreate_Fail() {
	bodyRequest := model.AssetLocation{
		Id: "1",
	}

	loc.mock.ExpectExec("INSERT INTO asset_location").WithArgs(bodyRequest.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := loc.repo.Create(bodyRequest)
	assert.Error(loc.T(), err)
}

func (loc *AssetLocationRepositorySuite) TestList_Success() {
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("1", "Location 1").
		AddRow("2", "Location 2")

	loc.mock.ExpectQuery("SELECT id, name FROM asset_location").WillReturnRows(rows)

	result, err := loc.repo.List()
	assert.NoError(loc.T(), err)
	assert.Len(loc.T(), result, 2)
}

func (loc *AssetLocationRepositorySuite) TestList_Fail() {
	loc.mock.ExpectQuery("SELECT id, name FROM asset_location").WillReturnError(sql.ErrNoRows)

	result, err := loc.repo.List()
	assert.Error(loc.T(), err)
	assert.Len(loc.T(), result, 0)
}

func (loc *AssetLocationRepositorySuite) TestListScan_Fail() {
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("1").
		AddRow("2")

	loc.mock.ExpectQuery("SELECT id, name FROM asset_location").WillReturnRows(rows)

	_, err := loc.repo.List()
	assert.Error(loc.T(), err)
}

func (loc *AssetLocationRepositorySuite) TestGet_Success() {
	id := "1"
	row := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("1", "Location 1")

	loc.mock.ExpectQuery("SELECT id, name FROM asset_location WHERE id=?").WithArgs(id).WillReturnRows(row)

	result, err := loc.repo.Get(id)
	assert.NoError(loc.T(), err)
	assert.Equal(loc.T(), id, result.Id)
}

func (loc *AssetLocationRepositorySuite) TestGet_Fail() {
	id := "1"

	loc.mock.ExpectQuery("SELECT id, name FROM asset_location WHERE id=?").WithArgs(id).WillReturnError(sql.ErrNoRows)

	result, err := loc.repo.Get(id)
	assert.NoError(loc.T(), err)
	assert.Equal(loc.T(), model.AssetLocation{}, result)
}

func (loc *AssetLocationRepositorySuite) TestUpdate_Success() {
	bodyRequest := model.AssetLocation{
		Id:   "1",
		Name: "Location 1",
	}

	loc.mock.ExpectExec("UPDATE asset_location").WithArgs(bodyRequest.Id, bodyRequest.Name).WillReturnResult(sqlmock.NewResult(1, 1))

	err := loc.repo.Update(bodyRequest)
	assert.NoError(loc.T(), err)
}

func (loc *AssetLocationRepositorySuite) TestUpdate_Fail() {
	bodyRequest := model.AssetLocation{
		Id: "1",
	}

	loc.mock.ExpectExec("UPDATE asset_location").WithArgs(bodyRequest.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := loc.repo.Update(bodyRequest)
	assert.Error(loc.T(), err)
}

func (loc *AssetLocationRepositorySuite) TestDelete_Success() {
	id := "1"

	loc.mock.ExpectExec("DELETE FROM asset_location WHERE id=?").WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := loc.repo.Delete(id)
	assert.NoError(loc.T(), err)
}

func (loc *AssetLocationRepositorySuite) TestDelete_Fail() {
	id := "1"

	loc.mock.ExpectExec("DELETE FROM vendors WHERE id=?").WithArgs(id).WillReturnError(sql.ErrNoRows)

	err := loc.repo.Delete(id)
	assert.Error(loc.T(), err)
}

func TestAssetLocationRepositorySuite(t *testing.T) {
	suite.Run(t, new(AssetLocationRepositorySuite))
}
