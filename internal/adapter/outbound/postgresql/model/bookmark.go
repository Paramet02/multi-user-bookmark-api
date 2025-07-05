package model

import (
	"time"
)

// Bookmark represents a user's bookmark in the system.
// This struct is used to map the bookmark entity to the database.
type BookmarkModel struct {
	ID          int       		`gorm:"primaryKey"`
	UserID      int       		`gorm:"not null"`
	URLID	   string     		`gorm:"type:varchar(255);uniqueIndex"`
	Title       string    		`gorm:"type:varchar(255)"`
	Description string    		`gorm:"type:text"`
	CreatedAt   time.Time 		`gorm:"autoCreateTime"`
	Tags        []TagModel    	`gorm:"many2many:bookmark_tags;"`
}

// TableName overrides table name
func (BookmarkModel) TableName() string {
	return "bookmarks"
}