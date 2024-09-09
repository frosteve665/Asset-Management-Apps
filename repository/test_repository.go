package repository

import (
	"database/sql"
)

type TestRepo interface {
	Test() error
}

type testRepo struct {
	db *sql.DB
}

func (c *testRepo) Test() error {
	return nil
}

func NewTestRepository(db *sql.DB) TestRepo {
	return &testRepo{
		db: db,
	}
}
