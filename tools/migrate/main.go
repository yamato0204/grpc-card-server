package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
	mmysql "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/google/martian/v3/log"
	"github.com/yamato0204/grpc-card-server/internal/infra/mysql"
)

const (
	migrateDir = "db"
)

func main() {

	// master
	err := migrateMysqlDB(&mysql.Config{
		Addr:     os.Getenv("MYSQL_MASTER_ADDR"),
		Protocol: os.Getenv("MYSQL_MASTER_PROTOCOL"),
		User:     os.Getenv("MYSQL_MASTER_USER"),
		Password: os.Getenv("MYSQL_MASTER_PASSWORD"),
		DB:       os.Getenv("MYSQL_MASTER_DB"),
	})
	if err != nil {
		log.Errorf("failed to migrate master: %v", err)
		os.Exit(1)
	}

	// user
	err = migrateMysqlDB(&mysql.Config{
		Addr:     os.Getenv("MYSQL_SHARD_ADDR"),
		Protocol: os.Getenv("MYSQL_SHARD_PROTOCOL"),
		User:     os.Getenv("MYSQL_SHARD_USER"),
		Password: os.Getenv("MYSQL_SHARD_PASSWORD"),
		DB:       os.Getenv("MYSQL_SHARD_DB"),
	})
	if err != nil {
		log.Errorf("failed to migrate user: %v", err)
		os.Exit(1)
	}
}

func migrateMysqlDB(cfg *mysql.Config) error {
	db, err := mysql.New(cfg)
	if err != nil {
		return err
	}

	fmt.Println("db is created")
	fmt.Println(db.Driver())

	driver, err := mmysql.WithInstance(db, &mmysql.Config{})
	if err != nil {
		log.Errorf("failed to create driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://./%s/%s", migrateDir, cfg.DB),
		"mysql",
		driver,
	)
	if err != nil {
		log.Errorf("failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Errorf("failed to migrate: %v", err)
	}

	log.Infof("migrate success")
	return nil
}
