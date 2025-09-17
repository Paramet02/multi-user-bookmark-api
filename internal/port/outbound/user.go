package outbound

import (
	"github.com/Paramet02/multi-user-bookmark-api/internal/domain"
	"context"
)

type UserRepository interface {
	// interface for user repository
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id int) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User , error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, user *domain.User) error
}