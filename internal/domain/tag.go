package domain

import "time"

// entity for Tag
// not have no tags or gorm annotations
type Tag struct {
	// Entity fields
	ID    int
	Name  string
	Color string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time // pointer to time.Time to allow null value
}
