package security

import (
    "context"
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound/security"
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/outbound/security/model"
)

type jwtManager struct {
    secretKey     string
    tokenDuration time.Duration
}

type contextKey string

const userIDKey contextKey = "user_id"
const userRole contextKey = "role"

func NewJWTManager(secretKey string, tokenDuration time.Duration) security.JWTManager {
    return &jwtManager{
        secretKey:     secretKey,
        tokenDuration: tokenDuration,
    }
}

func (j *jwtManager) GenerateToken(id int, role string) (string, error) {
    claims := &model.Claims{
        UserID: id,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenDuration)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(j.secretKey))
}

func (j *jwtManager) ValidateToken(tokenStr string) (int, string, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(j.secretKey), nil
    })

    if err != nil {
        return 0, "", err
    }

    claims, ok := token.Claims.(*model.Claims)
    if !ok || !token.Valid {
        return 0, "" ,errors.New("invalid token")
    }

    if claims.ExpiresAt.Time.Before(time.Now()) {
        return 0, "",errors.New("token expired")
    }

    return claims.UserID, claims.Role ,nil
}

func (j *jwtManager) InjectUserID(ctx context.Context, userID int, role string) context.Context {
    ctx = context.WithValue(ctx, userIDKey, userID)
    ctx = context.WithValue(ctx, userRole , role)
    return ctx
}

func (j *jwtManager) ExtractUserID(ctx context.Context) (int, string,error) {
    val := ctx.Value(userIDKey)
    if val == nil {
        return 0, "", errors.New("user ID not found in context")
    }

    role := ctx.Value(userRole)
    if role == nil {
        return 0, "" ,errors.New("role not match")
    }

    userID, ok := val.(int)
    if !ok {
        return 0, "" ,errors.New("user ID is not an int")
    }

    userRole , ok := role.(string)
    if !ok {
        return 0, "" ,errors.New("role is not an str")
    }
    return userID, userRole,nil
}
