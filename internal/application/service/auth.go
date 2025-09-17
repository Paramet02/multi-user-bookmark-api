package service

import (
	"context"
	"errors"
	"strings"

	"github.com/Paramet02/multi-user-bookmark-api/internal/domain"
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/inbound"
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound"
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound/security"
)

type authService struct {
	userRepo outbound.UserRepository
	hasher   security.PasswordHasher
	jwt 	 security.JWTManager
}

func NewAuthService(userRepo outbound.UserRepository, hasher security.PasswordHasher , jwt security.JWTManager) inbound.AuthService {
	return &authService{userRepo: userRepo , hasher: hasher , jwt: jwt}
}

func (a *authService) Login(ctx context.Context, identifier, password string) (string, error) {
    var user *domain.User
    var err error

    // ตรวจสอบ identifier ว่าเป็น email หรือ username
    if strings.Contains(identifier, "@") {
        user, err = a.userRepo.GetByEmail(ctx, identifier)
    } else {
        user, err = a.userRepo.GetByUsername(ctx, identifier)
    }

    if err != nil {
        return "", err
    }
    if user == nil {
        return "", errors.New("invalid email/username or password")
    }

    // ตรวจสอบรหัสผ่าน
    err = a.hasher.Compare(password, user.PasswordHash)
    if err != nil {
        return "", errors.New("invalid email or password")
    }

    token, err := a.jwt.GenerateToken(user.ID, user.Role)
    if err != nil {
        return "", err
    }

    return token, nil
}
