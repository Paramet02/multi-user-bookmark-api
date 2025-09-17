package mapper

import (
	"github.com/Paramet02/multi-user-bookmark-api/internal/domain"
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/outbound/postgresql/model"

	"gorm.io/gorm"
	"time"
)	

// ------------------------------------ Model Mapper Functions ---------------------------------- //
// -------------------- Methods to convert between domain User and UserModel -------------------- //

// toUserModel converts a domain User to a UserModel for database storage.
// domain.User → model.UserModel
// send return to database layer
func ToUserModel(u *domain.User) *model.UserModel {
	return &model.UserModel{
		Email:        u.Email,
		Username: 	  u.Username,
		PasswordHash: u.PasswordHash,
		Role:         u.Role,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		DeletedAt:    toGormDeletedAt(u.DeletedAt),
	}
}

// toDomainUser converts a UserModel from the database to a domain User.
// model.UserModel → domain.User
// send return to business logic layer
func ToDomainUser(m *model.UserModel) *domain.User {
	return &domain.User{
		Email:        m.Email,
		Username: 	  m.Username,
		PasswordHash: m.PasswordHash,
		Role:         m.Role,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		DeletedAt:    toTimePtr(m.DeletedAt),
	}
}

// help *time.Time to gorm.DeletedAt
func toGormDeletedAt(t *time.Time) gorm.DeletedAt {
	if t != nil {
		return gorm.DeletedAt{Time: *t, Valid: true}
	}
	return gorm.DeletedAt{}
}

// ช่วยแปลง gorm.DeletedAt เป็น *time.Time
func toTimePtr(d gorm.DeletedAt) *time.Time {
	if d.Valid {
		return &d.Time
	}
	return nil
}