package manager

import (
	"asetku-bukan-asetmu/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type InfraManager interface {
	Connection() *sql.DB
}

type infraManager struct {
	db  *sql.DB
	cfg *config.Config
}

func (i *infraManager) initDb() error {
	var dbConf = i.cfg.DBConfig
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		dbConf.Password,
		dbConf.Name)

	db, err := sql.Open(i.cfg.Driver, dataSource)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	i.db = db
	return nil
}

func (i *infraManager) Connection() *sql.DB {
	return i.db
}

func NewInfraManager(configParam *config.Config) (InfraManager, error) {
	infra := &infraManager{
		cfg: configParam,
	}

	err := infra.initDb()
	if err != nil {
		return nil, err
	}

	return infra, nil
}
