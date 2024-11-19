package master

import (
	"context"

	"github.com/yamato0204/grpc-card-server/internal/domain/repository/card"
	"github.com/yamato0204/grpc-card-server/internal/domain/repository/character"
	pb "github.com/yamato0204/grpc-card-server/pkg/pb/service"
)

type handler struct {
	cardRepository      card.Repository
	characterRepository character.Repository
}

func New(
	cardRepository card.Repository,
	characterRepository character.Repository,
) pb.MasterServiceServer {
	return &handler{
		cardRepository:      cardRepository,
		characterRepository: characterRepository,
	}
}

func (h *handler) GetAll(ctx context.Context, _ *pb.Empty) (*pb.GetAllResponse, error) {
	cards, err := h.cardRepository.SelectAll(ctx)
	if err != nil {
		return nil, err
	}
	characters, err := h.characterRepository.SelectAll(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetAllResponse{
		Cards:      toCardMasterList(cards),
		Characters: toCharacterMasterList(characters),
	}, nil
}

func (h *handler) GetCard(ctx context.Context, _ *pb.Empty) (*pb.GetCardResponse, error) {
	cards, err := h.cardRepository.SelectAll(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetCardResponse{
		Cards: toCardMasterList(cards),
	}, nil
}

func (h *handler) GetCharacter(ctx context.Context, _ *pb.Empty) (*pb.GetCharacterResponse, error) {
	characters, err := h.characterRepository.SelectAll(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetCharacterResponse{
		Characters: toCharacterMasterList(characters),
	}, nil
}
