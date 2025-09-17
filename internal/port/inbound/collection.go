package inbound

import (
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/inbound/http/response/model"
	"context"
)

// port for collection service
type CollectionService interface {
	// interface for collection service
	CreateCollection(ctx context.Context, userID int, name string) (*model.CollectionResponse, error)
	GetCollectionByID(ctx context.Context, userid int ,id int) (*model.CollectionResponse, error)
	GetCollectionsByUserID(ctx context.Context, userID int) ([]*model.CollectionResponse, error)
	UpdateCollection(ctx context.Context, id int, name string) (*model.CollectionResponse, error)	
	DeleteCollection(ctx context.Context, id int) error
}