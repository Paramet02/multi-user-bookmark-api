package domain

import (
	"time"
)

// entity for bookmark
// not have no tags or gorm annotations
type Bookmark struct {
	// Entity fields
	ID           int 
	OwnerID      int 
	CollectionID *int 
	URLID        string
	Title        string
	Description  string 
	Note         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
