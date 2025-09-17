package repository

// repository not handling errors sand error to business logic layer handling
import (
	"time"
	"context"
	"errors"

	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/outbound/postgresql/mapper"
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/outbound/postgresql/model"
	"github.com/Paramet02/multi-user-bookmark-api/internal/domain"
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound"
	"gorm.io/gorm"
)

// adapter implementing outbound.UserRepository interface and using gorm to interact with the database.
type userRepositoryDB struct {
	db *gorm.DB
}

// factory function to create a new instance of userRepositoryDB.
func NewUserRepositoryDB(db *gorm.DB) outbound.UserRepository {
	return &userRepositoryDB{db: db}
}

// method base on outbound.UserRepository interface
func (r *userRepositoryDB) Create(ctx context.Context, user *domain.User) error {
	userModel := mapper.ToUserModel(user)
	userModel.CreatedAt = time.Now()
	userModel.UpdatedAt = time.Now()

	if err := r.db.WithContext(ctx).Create(userModel).Error; err != nil {
		return err
	}

	// Copy auto-increment ID จาก DB กลับไป domain.User
	user.ID = userModel.ID

	return nil
}

func (r *userRepositoryDB) GetByID(ctx context.Context , id int) (*domain.User, error) {
	var model model.UserModel
	if err := r.db.WithContext(ctx).First(&model , id).Error; err != nil {
		return nil , err
	}
	return mapper.ToDomainUser(&model), nil
}

func (r *userRepositoryDB) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
    var user model.UserModel
    err := r.db.WithContext(ctx).
        Where("email = ?", email).
        First(&user).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err 
    }
	users := mapper.ToDomainUser(&user)

	return users, nil
}


func (r *userRepositoryDB) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
    var user model.UserModel
    err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return mapper.ToDomainUser(&user), nil
}

func (r *userRepositoryDB) Update(ctx context.Context , user *domain.User) error {
	userModel := mapper.ToUserModel(user)
	userModel.UpdatedAt = time.Now() // Update the UpdatedAt field to the current time

	if err := r.db.WithContext(ctx).Save(userModel).Error; err != nil {
		return err
	}
	
	return nil
}

func (r *userRepositoryDB) Delete(ctx context.Context , id int) error {
	var model model.UserModel
	
	result := r.db.WithContext(ctx).Delete(&model , id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }
	
	return nil
}

