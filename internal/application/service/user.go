package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/inbound/http/response/mapper"
	"github.com/Paramet02/multi-user-bookmark-api/internal/adapter/inbound/http/response/model"

	"github.com/Paramet02/multi-user-bookmark-api/internal/port/inbound"
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound"
	"github.com/Paramet02/multi-user-bookmark-api/internal/port/outbound/security"
	"github.com/Paramet02/multi-user-bookmark-api/internal/domain"
	"context"
)

// userService implements the inbound.UserService interface 
type userService struct {
	// userRepo and hasher , Policy are dependencies injected into the userService (DIP - Dependency Inversion Principle)
	userRepo  	outbound.UserRepository
	hasher    	security.PasswordHasher
	policy 		security.Policy
}

// factory function to create a new instance of userService
// It takes userRepo and hasher as parameters, which are used to interact with the user
func NewUserService(userRepo outbound.UserRepository, hasher security.PasswordHasher, policy security.Policy) inbound.UserService {
	return &userService{
		userRepo: userRepo,
		hasher:   hasher,
		policy:   policy,
		
	}
}

func (s *userService) RegisterUser(ctx context.Context , email, username, password string) (*model.UserResponse, error) {
	
	// Validate email and password
	if err := s.policy.EmailValidate(email); err != nil {
		return nil, err
	}

	// check email
	existingUser, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}	

	// validate username 
	if err := s.policy.UsernameValidate(username); err != nil {
		return nil , err
	}

	// check username 
	existingUsername , err := s.userRepo.GetByUsername(ctx ,username)
	if err != nil {
		return nil , fmt.Errorf("faild to check user ")
	}

	if existingUsername != nil {
		return nil , errors.New("user with this username already exists")
	}
	
	// policy for password
	if err := s.policy.PasswordValidate(password); err != nil {
		return nil, err
	}

	// Hash the password
	hashedPassword, err := s.hasher.Hash(password)
	if err != nil {
		return nil , errors.New("failed to hash password : " + err.Error())
	}

	// Create a new user
	newUser := &domain.User{
		Email: email,
		Username: username,
		PasswordHash: hashedPassword,
		Role: "user", // default role
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	// Save to DB
	if err := s.userRepo.Create(ctx, newUser); err != nil {
		return nil, errors.New("failed to create user: " + err.Error())
	}

	// Map to response
	userResponse := mapper.ToUserResponse(newUser)
	userResponse.ID = newUser.ID 

	return userResponse, nil
}


func (s *userService) GetUserByID(ctx context.Context , id int) (*model.UserResponse, error) {
	if id <= 0 {
		return nil , errors.New("invalid user ID")
	}

	user , err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("failed to get user by ID: " + err.Error())
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	userResponse := mapper.ToUserResponse(user)
	if userResponse == nil {
		return nil, errors.New("failed to map user to response")
	}

	return userResponse , nil
}

func (s *userService) GetUserByEmail(ctx context.Context , email string) (*model.UserResponse, error) {
	if email == "" {
		return nil , errors.New("email can't be empty")
	}

	user , err := s.userRepo.GetByEmail(ctx, email)

	if err != nil {
		return nil , errors.New("failed to get user by email : " + err.Error())
	}

	if user == nil {
		return nil , errors.New("user not found")
	}
	userReponse := mapper.ToUserResponse(user)

	if userReponse == nil { 
		return nil , errors.New("failed to map user to response")
	}

	return userReponse , nil
}

func (s *userService) GetUserByUsername(ctx context.Context , username string) (*model.UserResponse, error) {
	if username == "" {
		return  nil , errors.New("email can't be empty")
	}

	user , err := s.userRepo.GetByUsername(ctx ,username)

	if err != nil {
		return nil , errors.New("failed to get user by username : " + err.Error())
	}

	if user == nil {
		return nil , errors.New("user not found")
	}

	userReponse := mapper.ToUserResponse(user)

	if userReponse == nil { 
		return nil , errors.New("failed to map user to response")
	}

	return userReponse , nil
}

func (s *userService) UpdateUser(ctx context.Context, id int, email, username, password string) (*model.UserResponse, error) {
	// 1. Validate user ID
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}

	// 2. Get existing user by ID
	existingUser, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("user not found: " + err.Error())
	}
	if existingUser == nil {
		return nil, errors.New("user not found")
	}

	// 3. Update email if provided
	if email != "" && email != existingUser.Email {
		if err := s.policy.EmailValidate(email); err != nil {
			return nil, errors.New("invalid email: " + err.Error())
		}

		userWithSameEmail, err := s.userRepo.GetByEmail(ctx, email)
		if err == nil && userWithSameEmail != nil && userWithSameEmail.ID != id {
			return nil, errors.New("email already in use by another user")
		}

		existingUser.Email = email
	}

	// 4. Update username if provided
	if username != "" && username != existingUser.Username {
		if err := s.policy.UsernameValidate(username); err != nil {
			return nil, errors.New("invalid username: " + err.Error())
		}

		userWithSameUsername, err := s.userRepo.GetByUsername(ctx, username)
		if err == nil && userWithSameUsername != nil && userWithSameUsername.ID != id {
			return nil, errors.New("username already in use by another user")
		}

		existingUser.Username = username
	}

	// 5. Update password if provided
	if password != "" {
		if err := s.policy.PasswordValidate(password); err != nil {
			return nil, errors.New("password is too weak: " + err.Error())
		}

		hashedPassword, err := s.hasher.Hash(password)
		if err != nil {
			return nil, errors.New("failed to hash password: " + err.Error())
		}

		existingUser.PasswordHash = hashedPassword
	}

	// 6. Update time
	existingUser.UpdatedAt = time.Now()

	// 7. Save to DB
	if err := s.userRepo.Update(ctx, existingUser); err != nil {
		return nil, errors.New("failed to update user: " + err.Error())
	}

	// 8. Return mapped response
	return mapper.ToUserResponse(existingUser), nil
}

func (s *userService) DeleteUser(ctx context.Context , id int) error {
	// user id
	if id <= 0 {
		return errors.New("user id under zero")  
	}

	// get user form database
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("failed to get user by ID: " + err.Error())
	}

	if user == nil {
		return errors.New("user not found")
	}

	// delete user from database
	if err := s.userRepo.Delete(ctx, id); err != nil {
		return errors.New("failed to delete user: " + err.Error())
	}

	// return nil if user deleted successfully
	return nil
}