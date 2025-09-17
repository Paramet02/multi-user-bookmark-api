package outbound

import (
	"github.com/Paramet02/multi-user-bookmark-api/internal/domain"
	"context"
)

type CollectionRepository interface {
	// interface for collection repository
	InsertCollection(ctx context.Context, collection *domain.Collection) error
	FindCollectionID(ctx context.Context, id int) (*domain.Collection, error)
	FindCollectionUserID(ctx context.Context, userID int) ([]*domain.Collection, error)
	FindCollectionUserIDAndName(ctx context.Context, userID int, name string) (*domain.Collection, error)
	DeleteCollection(ctx context.Context, id int) error
	UpdateCollection(ctx context.Context, collection *domain.Collection) error
}