package domain

import "time"

type BookmarkPermission struct {
	// Entity fields
	BookmarkID int
	UserID     int
	Role       string // owner, editor, viewer
	CreatedAt  time.Time
}
