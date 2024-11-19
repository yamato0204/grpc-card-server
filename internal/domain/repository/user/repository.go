package user

import (
	"context"

	"github.com/yamato0204/grpc-card-server/internal/domain/entity/shard"
)

type Repository interface {
	SelectAll(ctx context.Context) (shard.UserSlice, error)
	FindByID(ctx context.Context, userID string) (*shard.User, error)
	Insert(ctx context.Context, user *shard.User) error
	Update(ctx context.Context, user *shard.User) error
}
