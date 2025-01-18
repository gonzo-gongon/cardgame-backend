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

type CardUsecase struct {
	cardRepository repository.CardRepository
}

func (u *CardUsecase) GetCards(cardIDs []model.UUID[model.Card]) ([]model.Card, error) {
	user, err := u.cardRepository.GetCards(cardIDs)
	if err != nil {
		return nil, &CardUsecaseGetCardsError{
			cause: err,
		}
	}

	return user, nil
}

func NewCardUsecase(
	cardRepository repository.CardRepository,
) (*CardUsecase, error) {
	return &CardUsecase{
		cardRepository: cardRepository,
	}, nil
}
