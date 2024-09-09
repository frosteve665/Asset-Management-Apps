package repository_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/repository"
)

type VendorRepositorySuite struct {
	suite.Suite
	db         *sql.DB
	mock       sqlmock.Sqlmock
	repository repository.VendorRepository
}

func (s *VendorRepositorySuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(s.T(), err)

	s.db = db
	s.mock = mock
	s.repository = repository.NewVendorRepository(db)
}

func (s *VendorRepositorySuite) TearDownTest() {
	s.db.Close()
}

func (s *VendorRepositorySuite) TestCreateSuccess() {
	payload := model.Vendor{
		Id:      "1",
		Name:    "Vendor 1",
		Address: "Jl. Vendor 1",
		Phone:   "08123456789",
	}

	s.mock.ExpectExec("INSERT INTO vendors").WithArgs(payload.Id, payload.Name, payload.Address, payload.Phone).WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.Create(payload)
	assert.NoError(s.T(), err)
}

func (s *VendorRepositorySuite) TestCreateFail() {
	payload := model.Vendor{
		Id:      "1",
		Name:    "Vendor 1",
		Address: "Jl. Vendor 1",
	}

	s.mock.ExpectExec("INSERT INTO vendors").WithArgs(payload.Id, payload.Name, payload.Address).WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.Create(payload)
	assert.Error(s.T(), err)
}

func (s *VendorRepositorySuite) TestListSuccess() {
	rows := sqlmock.NewRows([]string{"id", "name", "address", "phone"}).
		AddRow("1", "Vendor 1", "Jl. Vendor 1", "08123456789").
		AddRow("2", "Vendor 2", "Jl. Vendor 2", "08123456789")

	s.mock.ExpectQuery("SELECT id, name, address, phone FROM vendors").WillReturnRows(rows)

	result, err := s.repository.List()
	assert.NoError(s.T(), err)
	assert.Len(s.T(), result, 2)
}

func (s *VendorRepositorySuite) TestListFail() {
	s.mock.ExpectQuery("SELECT id, name, address, phone FROM vendors").WillReturnError(sql.ErrNoRows)

	result, err := s.repository.List()
	assert.Error(s.T(), err)
	assert.Len(s.T(), result, 0)
}

func (s *VendorRepositorySuite) TestGetSuccess() {
	id := "1"
	row := sqlmock.NewRows([]string{"id", "name", "address", "phone"}).
		AddRow("1", "Vendor 1", "Jl. Vendor 1", "08123456789")

	s.mock.ExpectQuery("SELECT id, name, address, phone FROM vendors WHERE id = ?").WithArgs(id).WillReturnRows(row)

	result, err := s.repository.Get(id)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), id, result.Id)
}

func (s *VendorRepositorySuite) TestGetFail() {
	id := "1"

	s.mock.ExpectQuery("SELECT id, name, address, phone FROM vendors WHERE id = ?").WithArgs(id).WillReturnError(sql.ErrNoRows)

	result, err := s.repository.Get(id)
	assert.Error(s.T(), err)
	assert.Equal(s.T(), model.Vendor{}, result)
}

func (s *VendorRepositorySuite) TestUpdateSuccess() {
	payload := model.Vendor{
		Id:      "1",
		Name:    "Vendor 1 Updated",
		Address: "Jl. Vendor 1 Updated",
		Phone:   "08123456789",
	}

	s.mock.ExpectExec("UPDATE vendors").WithArgs(payload.Name, payload.Address, payload.Phone, payload.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.Update(payload)
	assert.NoError(s.T(), err)
}

func (s *VendorRepositorySuite) TestUpdateFail() {
	payload := model.Vendor{
		Id:      "1",
		Name:    "Vendor 1 Updated",
		Address: "Jl. Vendor 1 Updated",
	}

	s.mock.ExpectExec("UPDATE vendors").WithArgs(payload.Name, payload.Address, payload.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.Update(payload)
	assert.Error(s.T(), err)
}

func (s *VendorRepositorySuite) TestDeleteSuccess() {
	id := "1"

	s.mock.ExpectExec("DELETE FROM vendors WHERE id = ?").WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repository.Delete(id)
	assert.NoError(s.T(), err)
}

func (s *VendorRepositorySuite) TestDeleteFail() {
	id := "1"

	s.mock.ExpectExec("DELETE FROM vendors WHERE id = ?").WithArgs(id).WillReturnError(sql.ErrNoRows)

	err := s.repository.Delete(id)
	assert.Error(s.T(), err)
}

func TestVendorRepositorySuite(t *testing.T) {
	suite.Run(t, new(VendorRepositorySuite))
}
