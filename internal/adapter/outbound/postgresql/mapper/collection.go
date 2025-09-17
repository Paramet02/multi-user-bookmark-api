package mapper

import (
	"github.com/Paramet02/multi-user-bookmark-api/internal/domain"
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/outbound/postgresql/model"

)	

// ------------------------------------ Model Mapper Functions ---------------------------------- //
// -------------------- Methods to convert between domain Collection and CollectionModel -------------------- //
// toCollectionModel converts a domain Collection to a CollectionModel for database storage.
// domain.Collection → model.CollectionModel
// send return to database layer
func ToCollectionModel(c *domain.Collection) *model.CollectionModel {
	return &model.CollectionModel{
		UserID:    c.UserID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: toGormDeletedAt(c.DeletedAt),
	}
}

// toDomainCollection converts a CollectionModel from the database to a domain Collection.
// model.CollectionModel → domain.Collection
// send return to business logic layer
func ToDomainCollection(m *model.CollectionModel) *domain.Collection {
	return &domain.Collection{
		ID:        m.ID,
		UserID:    m.UserID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: toTimePtr(m.DeletedAt),
	}
}