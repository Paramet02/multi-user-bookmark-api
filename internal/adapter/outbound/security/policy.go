package security

import (
	"errors"
	"regexp"
	"strings"

	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound/security"
	"github.com/wagslane/go-password-validator"
)

// implementation of PasswordPolicy interface
type Policy struct {
	// minLength is the minimum length of the password
	minEntropy float64
	pattern *regexp.Regexp
	patternname *regexp.Regexp
}

var (
	// blacklist substrings
	blackListSub = []string{
		"root",
		"admin",
		"administrator",
		"system",
		"guest",
		"support",
	}
)

// factory function to create a new instance of PasswordPolicy
func NewPolicy(minEntropy float64) security.Policy {
	return &Policy{
		minEntropy: minEntropy,
		pattern: regexp.MustCompile(`^[a-zA-Z0-9]+@(gmail|hotmail)\.com$`),
		patternname: regexp.MustCompile(`^[a-zA-Z0-9_-]{3,16}$`),
	}
}

func (p *Policy) PasswordValidate(password string) error {
	if err := passwordvalidator.Validate(password , p.minEntropy); err != nil {
		return errors.New("password is too weak")
	}
		return nil
}

func (p *Policy) EmailValidate(email string) error {
	if email == "" {
		return errors.New("email empty")
	}

    if !p.pattern.MatchString(email) {
        return errors.New("email does not match policy")
    }
    return nil
}

func (p *Policy) UsernameValidate(username string) error {
	if username == "" {
		return errors.New("username: empty")
	}

	if !p.patternname.MatchString(username) {
		return errors.New("username does not match policy")
	}

	for _, banned := range blackListSub {
		if strings.Contains(username , banned) {
			return errors.New("username: contains reserved word")
		}
	}
	return nil
}

