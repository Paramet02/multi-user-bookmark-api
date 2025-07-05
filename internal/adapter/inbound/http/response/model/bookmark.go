package model

import "time"

// UserResponse represents the response structure for a user.
// It is used to send user data in API responses.
// This struct is used to map the User entity to a response format.
type BookmarkResponse struct {
	ID        int       		`json:"id"`
	URLID	   string     		`json:"url_id"`
	Title       string    		`json:"title"`
	Description string    		`json:"Description"`
	CreatedAt   time.Time 		`json:"CreatedAt"`
	Tags        []TagResponse 	`json:"tags"`
}

// method to get the current time
func GetCurrentTime() time.Time {
	return time.Now()
}
