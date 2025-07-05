package model

import "time"

// UserResponse represents the response structure for a user.
// It is used to send user data in API responses.
// This struct is used to map the User entity to a response format.
type UserResponse struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time	`json:"created_at"`
	UpdatedAt    time.Time	`json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"` // pointer to time.Time to allow null value
}
