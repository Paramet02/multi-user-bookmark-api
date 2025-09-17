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
func ToUserResponse(user *domain.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        user.ID,
		Username:  user.Username, 
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}


// ToDomainUser converts a UserResponse from the HTTP response to a domain User.
// response.UserResponse → domain.User
// This function is used to map the UserResponse back to the domain User entity.
func ToDomainUser(response *model.UserResponse) *domain.User {
	return &domain.User{
		Username:     response.Username,
		Email:        response.Email,
		CreatedAt:    response.CreatedAt,
		UpdatedAt:    response.UpdatedAt,
	}
}

func ToUserPublicResponse(user *model.UserResponse) *model.UserPublicResponse {
	return &model.UserPublicResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}