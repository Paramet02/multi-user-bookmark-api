package outbound

import (
	"github.com/Paramet02/multi-user-bookmark-api/internal/domain"
)

type UserRepository interface {
	// interface for user repository
	Create(user *domain.User) (int, error)
	GetByID(id int) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Delete(id int) error
	Update(user *domain.User) error
}