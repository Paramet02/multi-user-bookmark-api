package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/inbound/http/response/mapper"
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/inbound/http/response/model"
	"github.com/Paramet02/multi-user-bookmark-api/internal/domain"

	"github.com/Paramet02/multi-user-bookmark-api/internal/port/inbound"
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound"
)

// collectionService struct implements the inbound.CollectionService interface
type collectionService struct {
	collectionRepo outbound.CollectionRepository
}

// NewCollectionService is a factory function to create a new instance of collectionService
func NewCollectionService(collectionRepo outbound.CollectionRepository) inbound.CollectionService {
	return &collectionService{
		collectionRepo: collectionRepo,
	}
}

// CreateCollection creates a new collection for a user
func (s *collectionService) CreateCollection(ctx context.Context, userID int, name string) (*model.CollectionResponse , error) {
	// check if collection name already exists for th uesr
	if name == "" {
		return nil, errors.New("collection name can't be empty")
	}

	if len(name) > 100 {
		return nil , errors.New("collection name can't be more than 100 characters")
	}

	if existingCollection, _ := s.collectionRepo.FindCollectionUserIDAndName(ctx, userID, name); existingCollection != nil {
		return nil, errors.New("collection name already exists")
	}

	newCollection := &domain.Collection{
		UserID: userID,
		Name:  name,
	}

	if err := s.collectionRepo.InsertCollection(ctx, newCollection); err != nil {
		return nil, fmt.Errorf("failed to create collection: %w", err)
	}

	return mapper.ToCollectionResponse(newCollection) , nil
}

func (s *collectionService) GetCollectionByID(ctx context.Context, userid int ,id int) (*model.CollectionResponse, error) {
	
	return nil, nil
}

func (s *collectionService) GetCollectionsByUserID(ctx context.Context, userID int) ([]*model.CollectionResponse, error) {
	return nil, nil
}

func (s *collectionService) UpdateCollection(ctx context.Context, id int, name string) (*model.CollectionResponse, error) {
	return nil, nil
}

func (s *collectionService) DeleteCollection(ctx context.Context, id int) error {
	return nil
}
