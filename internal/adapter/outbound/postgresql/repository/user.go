package repository

// repository not handling errors sand error to business logic layer handling
import (
	"time"

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
func (r *userRepositoryDB) Create(user *domain.User) (int , error) {
	// convert domain.User to model.UserModel because gorm will use model.UserModel to create a new user in the database
	model := mapper.ToUserModel(user)
	if err := r.db.Create(model).Error; err != nil {
		return 0 , err
	}
	return model.ID , nil
}

func (r *userRepositoryDB) GetByID(id int) (*domain.User, error) {
	var model model.UserModel
	if err := r.db.First(&model , id).Error; err != nil {
		return nil , err
	}
	return mapper.ToDomainUser(&model), nil
}

func (r *userRepositoryDB) GetByEmail(email string) (*domain.User, error) {
	var model model.UserModel
	if err := r.db.Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}
	return mapper.ToDomainUser(&model), nil
}

func (r *userRepositoryDB) Update(user *domain.User) error {
	model := mapper.ToUserModel(user)
	model.UpdatedAt = time.Now() // Update the UpdatedAt field to the current time

	if err := r.db.Save(&model).Error; err != nil {
		return err
	}
	
	return nil
}

func (r *userRepositoryDB) Delete(id int) error {
	var model model.UserModel
	if err := r.db.First(&model, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&model).Error; err != nil {
		return err
	}

	return nil
}

