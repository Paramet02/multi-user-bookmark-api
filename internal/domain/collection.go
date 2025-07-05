package domain

import "time"

// Collection represents a collection of bookmarks.
// entity for Collection
// not have no tags or gorm annotations
type Collection struct {
	// Entity fields
	ID        int
	UserID    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}