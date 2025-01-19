package usecase

import (
	"fmt"
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/repository"
)

type CardUsecaseGetCardsError struct {
	cause error
}

func (e *CardUsecaseGetCardsError) Error() string {
	return fmt.Sprintf("card usecase get cards failed: %o", e.cause)
}

type CardUsecaseCreateCardError struct {
	cause error
}

func (e *CardUsecaseCreateCardError) Error() string {
	return fmt.Sprintf("card usecase create card failed: %o", e.cause)
}

type CardUsecase struct {
	cardRepository repository.CardRepository
}

func (u *CardUsecase) GetCards(cardIDs []model.UUID[model.Card]) ([]model.Card, error) {
	cards, err := u.cardRepository.GetCards(cardIDs)
	if err != nil {
		return nil, &CardUsecaseGetCardsError{
			cause: err,
		}
	}

	return cards, nil
}

func (u *CardUsecase) CreateCard(createCard model.CreateCard, createdBy *model.User) (*model.Card, error) {
	card, err := u.cardRepository.CreateCard(createCard, createdBy.ID)
	if err != nil {
		return nil, &CardUsecaseCreateCardError{
			cause: err,
		}
	}

	return card, nil
}

func NewCardUsecase(
	cardRepository repository.CardRepository,
) (*CardUsecase, error) {
	return &CardUsecase{
		cardRepository: cardRepository,
	}, nil
}
