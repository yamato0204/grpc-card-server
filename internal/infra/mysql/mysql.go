package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

	if err := tryPing(db, conf.DB, 5); err != nil {
		return nil, err
	}

	return db, nil
}

func tryPing(db *sql.DB, dbName string, retries int) error {
	var err error
	for i := 0; i < retries; i++ {
		err = db.Ping()
		if err == nil {
			fmt.Printf("Successfully connected to database: %s\n", dbName)
			return nil
		}
		fmt.Printf("Failed to connect to database %s. Retrying... (%d/%d)\n", dbName, i+1, retries)
		time.Sleep(2 * time.Second) // リトライ間隔（2秒）
	}
	return fmt.Errorf("could not connect to database %s after %d retries: %v", dbName, retries, err)
}
