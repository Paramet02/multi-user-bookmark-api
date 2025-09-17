package security

import (
	"fmt"

	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound/security"
	"golang.org/x/crypto/bcrypt"
)

type bcryptHasher struct{}

// factory function to create a new instance of BcryptHasher
func NewBcryptHasher() security.PasswordHasher {
	return &bcryptHasher{}
}

// hash password using bcrypt
func (b *bcryptHasher) Hash(password string) (string , error) {
	hashedpassword ,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "" , err
	}
	return string(hashedpassword) , nil
}

// ---------------- Bcrypt Hasher ----------------
func (b *bcryptHasher) Compare(password , hash string) error {
    fmt.Printf("DEBUG bcrypt.Compare -> hash='%s', len=%d, password='%s', len=%d\n", hash, len(hash), password, len(password))
	err := bcrypt.CompareHashAndPassword([]byte(hash) , []byte(password))
	if err != nil {
		return err
	}
	return nil
}