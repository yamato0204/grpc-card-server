package usercard

import (
	"context"

	"github.com/yamato0204/grpc-card-server/internal/domain/entity/shard"
)

type Repository interface {
	SelectByUserID(ctx context.Context, userID string) (shard.UserCardSlice, error)
	FindByCardID(ctx context.Context, userID string, cardID uint32) (*shard.UserCard, error)
	Insert(ctx context.Context, userCard *shard.UserCard) error
	Update(ctx context.Context, userCard *shard.UserCard) error
}
