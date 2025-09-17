package mapper

import (
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/inbound/http/response/model"
	"github.com/Paramet02/multi-user-bookmark-api/internal/domain"
)


// ------------------------------------ Model Mapper Functions ---------------------------------- //
// -------------------- Methods to convert between domain User and UserResponse -------------------- //
// ToUserResponse converts a domain User to a UserResponse for HTTP response.
// domain.User → response.UserResponse
// This function is used to map the User entity to a response format suitable for API responses.
func ToCollectionResponse(collection *domain.Collection) *model.CollectionResponse {
	return &model.CollectionResponse{
		ID:        collection.ID,
		UserID:    collection.UserID,
		Name:      collection.Name,
		CreatedAt: collection.CreatedAt,
		UpdatedAt: collection.UpdatedAt,
	}
}