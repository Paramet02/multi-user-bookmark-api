package model

// Tag represents a user's Tag in the system.
// This struct is used to map the Tag entity to the database.
type TagModel struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255)"`
}

// TableName overrides table name
func (TagModel) TableName() string {
	return "tags"
}