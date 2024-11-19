package card

import (
	"context"
	"log"

	"github.com/yamato0204/grpc-card-server/internal/domain/entity/master"
	"github.com/yamato0204/grpc-card-server/internal/domain/repository/card"
	"github.com/yamato0204/grpc-card-server/internal/infra/mysql"
)

type repository struct {
	masterDB *mysql.MasterDB
}

func New(db *mysql.MasterDB) card.Repository {
	return &repository{
		masterDB: db,
	}
}

func (repo *repository) SelectAll(ctx context.Context) (master.CardSlice, error) {
	records, err := master.Cards().All(ctx, repo.masterDB)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return records, nil
}
