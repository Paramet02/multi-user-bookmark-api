package repository

// repository not handling errors sand error to business logic layer handling
import (
	"context"

	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/outbound/postgresql/mapper"
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/outbound/postgresql/model"
	"github.com/Paramet02/multi-user-bookmark-api/internal/domain"
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound"
	"gorm.io/gorm"
)

// adapter implementing outbound.CollectionRepository interface and using gorm to interact with the database.
type collectionRepositoryDB struct {
	db *gorm.DB
}

// factory function to create a new instance of collectionRepositoryDB.
func NewCollectionRepositoryDB(db *gorm.DB) outbound.CollectionRepository {
	return &collectionRepositoryDB{db: db}
}

func (r *collectionRepositoryDB) InsertCollection(ctx context.Context, collection *domain.Collection) error {
	collectionModel := mapper.ToCollectionModel(collection)
	if err := r.db.WithContext(ctx).Create(collectionModel).Error; err != nil {
		return err
	}
	
	// Copy auto-increment ID จาก DB กลับไป domain.Collection
	collection.ID = collectionModel.ID
	
	return nil
}

func (r *collectionRepositoryDB) FindCollectionID(ctx context.Context, id int) (*domain.Collection, error) {
	var model model.CollectionModel
	if err := r.db.WithContext(ctx).First(&model , id).Error; err != nil {
		return nil , err
	}
	return mapper.ToDomainCollection(&model), nil
}

func (r *collectionRepositoryDB) FindCollectionUserID(ctx context.Context, userID int) ([]*domain.Collection, error) {
	var models []model.CollectionModel
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&models).Error; err != nil {
		return nil ,err
	}

	if len(models) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var collections []*domain.Collection
	for _, m := range models {
		collections = append(collections, mapper.ToDomainCollection(&m))
	}

	return collections, nil
}

func (r *collectionRepositoryDB) FindCollectionUserIDAndName(ctx context.Context, userID int, name string) (*domain.Collection, error) {
	var model model.CollectionModel
	if err := r.db.WithContext(ctx).Where("user_id = ? AND name = ?", userID, name).First(&model).Error; err != nil {
		return nil , err
	}

	return mapper.ToDomainCollection(&model), nil
}


func (r *collectionRepositoryDB) DeleteCollection(ctx context.Context, id int) error {
	var collection model.CollectionModel
	result := r.db.WithContext(ctx).Delete(&collection , id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }

	return nil
}

func (r *collectionRepositoryDB) UpdateCollection(ctx context.Context, collection *domain.Collection) error {
	collectionModel := mapper.ToCollectionModel(collection)
	if err := r.db.WithContext(ctx).Model(&model.CollectionModel{}).Where("id = ?", collection.ID).Updates(collectionModel).Error; err != nil { 
		return err
	}
	
	return nil
}

