package inbound

import (
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/inbound/http/response/model"
	"context"
)

// port for user service
type UserService interface {
	// interface for user service
	RegisterUser(ctx context.Context, email , username ,password string) (*model.UserResponse, error)
	GetUserByID(ctx context.Context, id int) (*model.UserResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*model.UserResponse, error)
	GetUserByUsername(ctx context.Context, username string) (*model.UserResponse, error)
	UpdateUser(ctx context.Context, id int, email , username , password string) (*model.UserResponse, error)
	DeleteUser(ctx context.Context, id int) error
}