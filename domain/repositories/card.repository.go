package repositories

import (
	"cards/domain/entities"
	"github.com/google/uuid"
)

type CardRepository interface {
	GetAllAsync() (*entities.Card, error)
	GetAsync(uuid.UUID) (*entities.Card, error)
	AddAsync(*entities.Card) (*entities.Card, error)
	UpdateAsync(uuid.UUID, *entities.Card) (*entities.Card, error)
	DeleteAsync(uuid.UUID) (bool, error)
}
