package inbound

import (
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/inbound/http/response/model"
)

// port for user service
type UserService interface {
	// interface for user service
	RegisterUser(email , password string) (*model.UserResponse, error)
	GetUserByID(id int) (*model.UserResponse, error)
	GetUserByEmail(email string) (*model.UserResponse, error)
	UpdateUser(user *model.UserResponse) (*model.UserResponse, error)
	DeleteUser(id int) error
}