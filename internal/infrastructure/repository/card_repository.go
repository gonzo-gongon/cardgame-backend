package repository

import (
	domainmodel "original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/gateway"
	"original-card-game-backend/internal/infrastructure/model"
	"original-card-game-backend/internal/infrastructure/value"
)

type CardRepository interface {
	GetCards(cardIDs []domainmodel.UUID[domainmodel.Card]) ([]domainmodel.Card, error)
}

type CardRepositoryImpl struct {
	databaseGateway gateway.DatabaseGateway
}

func (r *CardRepositoryImpl) GetCards(cardIDs []domainmodel.UUID[domainmodel.Card]) ([]domainmodel.Card, error) {
	conn, err := r.databaseGateway.Connect()
	if err != nil {
		return nil, err
	}

	ids := value.UUIDsFromDomain[domainmodel.Card, model.Card](cardIDs)

	var cards model.Cards
	if result := conn.Where("id IN ?", ids).Find(&cards); result.Error != nil {
		return nil, result.Error
	}

	return cards.Domain(), nil
}

//nolint:ireturn // DIのためのコードなので許容する
func NewCardRepository(
	databaseGateway gateway.DatabaseGateway,
) (CardRepository, error) {
	return &CardRepositoryImpl{
		databaseGateway: databaseGateway,
	}, nil
}
