package model

import "time"

// CollectionResponse represents the response structure for a collection.
type CollectionResponse struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"` 
	Name      string `json:"name"`
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}