package model

// UserResponse represents the response structure for a user.
// It is used to send user data in API responses.
// This struct is used to map the User entity to a response format.
type TagResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
}