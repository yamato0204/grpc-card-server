package character

import (
	"context"

	"github.com/yamato0204/grpc-card-server/internal/domain/entity/master"
)

type Repository interface {
	SelectAll(ctx context.Context) (master.CharacterSlice, error)
}
