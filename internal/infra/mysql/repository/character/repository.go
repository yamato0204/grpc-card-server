package character

import (
	"context"
	"log"

	"github.com/yamato0204/grpc-card-server/internal/domain/entity/master"
	"github.com/yamato0204/grpc-card-server/internal/domain/repository/character"
	"github.com/yamato0204/grpc-card-server/internal/infra/mysql"
)

type repository struct {
	masterDB *mysql.MasterDB
}

func New(db *mysql.MasterDB) character.Repository {
	return &repository{
		masterDB: db,
	}
}

func (repo *repository) SelectAll(ctx context.Context) (master.CharacterSlice, error) {
	records, err := master.Characters().All(ctx, repo.masterDB)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return records, nil
}
