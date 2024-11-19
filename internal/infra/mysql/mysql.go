package mysql

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Addr     string
	Protocol string
	User     string
	Password string
	DB       string
}

type MasterDB struct {
	*sql.DB
}

func NewMasterDB(cfg *Config) (*MasterDB, error) {
	db, err := New(cfg)
	if err != nil {
		return nil, err
	}

	return &MasterDB{
		DB: db,
	}, nil
}

type ShardDB struct {
	*sql.DB
}

func NewShardDB(cfg *Config) (*ShardDB, error) {
	db, err := New(cfg)
	if err != nil {
		return nil, err
	}

	return &ShardDB{
		DB: db,
	}, nil
}

func New(conf *Config) (*sql.DB, error) {
	driver := "mysql"
	dsn := fmt.Sprintf(
		"%s:%s@%s(%s)/%s?parseTime=true&multiStatements=true",
		conf.User, conf.Password, conf.Protocol, conf.Addr, conf.DB,
	)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
