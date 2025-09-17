package security

import (
	"context"
)

type PasswordHasher interface{
	Hash(password string) (string , error)
	Compare(password, hash string) error
}

type Policy interface{
	PasswordValidate(password string) error
	EmailValidate(email string) error
	UsernameValidate(username string) error
}

type JWTManager interface{
	GenerateToken(id int, role string) (string , error)
	ValidateToken(tokenStr string) (int , string , error)
	InjectUserID(ctx context.Context, userID int ,role string) context.Context
	ExtractUserID(ctx context.Context) (int, string,error)
}