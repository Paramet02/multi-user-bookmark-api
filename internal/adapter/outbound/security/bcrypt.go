package security

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound/security"
)

type BcryptHasher struct{}

// factory function to create a new instance of BcryptHasher
func NewBcryptHasher() security.PasswordHasher {
	return &BcryptHasher{}
}

// hash password using bcrypt
func (b *BcryptHasher) Hash(password string) (string , error) {
	hashedpassword ,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "" , err
	}
	return string(hashedpassword) , nil
}

func (b *BcryptHasher) Compare(password , hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash) , []byte(password))
	if err != nil {
		return err
	}
	return nil
}