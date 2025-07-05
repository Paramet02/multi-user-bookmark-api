package security

import (
	"errors"
	"github.com/wagslane/go-password-validator"
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound/security"
)

// implementation of PasswordPolicy interface
type PasswordPolicy struct {
	// minLength is the minimum length of the password
	minEntropy float64
}

// factory function to create a new instance of PasswordPolicy
func NewPasswordPolicy(minEntropy float64) security.PasswordPolicy {
	return &PasswordPolicy{
		minEntropy: minEntropy,
	}
}

func (p *PasswordPolicy) Validate(password string) error {
	if err := passwordvalidator.Validate(password , p.minEntropy); err != nil {
		return errors.New("password is too weak")
	}
		return nil
}