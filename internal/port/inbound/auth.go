package inbound

import (
	"context"
)

type AuthService interface{
	Login(ctx context.Context, identifier,password string) (string , error)
} 