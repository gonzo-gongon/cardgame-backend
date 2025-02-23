package repository

import (
	domainmodel "original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/gateway"
	"original-card-game-backend/internal/infrastructure/model"
	"original-card-game-backend/internal/infrastructure/value"
)

type CardRepository interface {
	GetCards(cardIDs []domainmodel.UUID[domainmodel.Card]) ([]domainmodel.Card, error)
	CreateCard(
		createCard domainmodel.CreateCard,
		createBy domainmodel.UUID[domainmodel.User],
	) (*domainmodel.Card, error)
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

func (r *CardRepositoryImpl) CreateCard(
	createCard domainmodel.CreateCard,
	createBy domainmodel.UUID[domainmodel.User],
) (*domainmodel.Card, error) {
	conn, err := r.databaseGateway.Connect()
	if err != nil {
		return nil, err
	}

	ccb := value.UUIDFromDomain[domainmodel.User, model.User](createBy)

	card := model.Card{
		Name:      createCard.Name,
		Text:      createCard.Text,
		CreatedBy: &ccb,
		UpdatedBy: &ccb,
	}

	if result := conn.Create(&card); result.Error != nil {
		return nil, result.Error
	}

	ret := card.Domain()

	return &ret, nil
}

//nolint:ireturn // DIのためのコードなので許容する
func NewCardRepository(
	databaseGateway gateway.DatabaseGateway,
) (CardRepository, error) {
	return &CardRepositoryImpl{
		databaseGateway: databaseGateway,
	}, nil
}
