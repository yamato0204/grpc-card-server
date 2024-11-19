package master

import (
	"github.com/yamato0204/grpc-card-server/internal/domain/entity/master"
	"github.com/yamato0204/grpc-card-server/pkg/pb/enums"
	pbmaster "github.com/yamato0204/grpc-card-server/pkg/pb/master"
)

func toCardMasterList(lst master.CardSlice) []*pbmaster.Card {
	result := make([]*pbmaster.Card, 0, len(lst))
	for _, v := range lst {
		result = append(result, &pbmaster.Card{
			Id:          v.ID,
			Name:        v.Name,
			RarityType:  enums.CardRarityType(v.RarityType),
			CharacterId: v.CharacterID,
		})
	}

	return result
}

func toCharacterMasterList(lst master.CharacterSlice) []*pbmaster.Character {
	result := make([]*pbmaster.Character, 0, len(lst))
	for _, v := range lst {
		result = append(result, &pbmaster.Character{
			Id:     v.ID,
			Name:   v.Name,
			Gender: enums.Gender(v.Gender),
			Age:    v.Age,
		})
	}

	return result
}
