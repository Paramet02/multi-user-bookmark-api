package application

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
)

// userService implements the inbound.UserService interface 
type userService struct {
	// userRepo and hasher , Policy are dependencies injected into the userService (DIP - Dependency Inversion Principle)
	userRepo  	outbound.UserRepository
	hasher    	security.PasswordHasher
	policy 		security.PasswordPolicy
}

// factory function to create a new instance of userService
// It takes userRepo and hasher as parameters, which are used to interact with the user
func NewUserService(userRepo outbound.UserRepository, hasher security.PasswordHasher, policy security.PasswordPolicy) inbound.UserService {
	return &userService{
		userRepo: userRepo,
		hasher:   hasher,
		policy:   policy,
		
	}
}

func (s *userService) RegisterUser(email, password string) (*model.UserResponse, error) {
	
	// Validate email and password
	if email == "" || password == "" {
		return nil, errors.New("email and password can't be empty")
	}

	// check email
	existingUser, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}
	
	// policy for password
	if err := s.policy.Validate(password); err != nil {
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
		PasswordHash: hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	// Save the user to the repository (database)
	userID , err := s.userRepo.Create(newUser)

	if err != nil {
		return nil, errors.New("failed to create user: " + err.Error())
	}

	// mapper to convert domain.User to model.UserResponse
	userResponse := mapper.ToUserResponse(newUser)
	if userResponse == nil {
		return nil, errors.New("failed to map user to response")
	}

	userResponse.ID = userID
	userResponse.CreatedAt = newUser.CreatedAt
	userResponse.UpdatedAt = newUser.UpdatedAt

	return userResponse, nil
}


func (s *userService) GetUserByID(id int) (*model.UserResponse, error) {
	if id <= 0 {
		return nil , errors.New("invalid user ID")
	}

	user , err := s.userRepo.GetByID(id)
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

func (s *userService) GetUserByEmail(email string) (*model.UserResponse, error) {
	if email == "" {
		return  nil , errors.New("email can't be empty")
	}

	user , err := s.userRepo.GetByEmail(email)

	if err != nil {
		return  nil , errors.New("failed to get user by email : " + err.Error())
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

func (s *userService) UpdateUser(userResp *model.UserResponse) (*model.UserResponse, error) {
	// 1. Validate input
	if userResp == nil || userResp.Email == "" {
		return nil, errors.New("user or email can't be empty")
	}

	// 2. Map to domain user
	inputUser := mapper.ToDomainUser(userResp)
	if inputUser == nil {
		return nil, errors.New("failed to map user to domain")
	}

	// 3. Get existing user by ID
	existingUser, err := s.userRepo.GetByID(inputUser.ID)
	if err != nil {
		return nil, errors.New("user not found: " + err.Error())
	}

	// 4. Check if email changed and is already used by another user
	if inputUser.Email != existingUser.Email {
		userWithSameEmail, err := s.userRepo.GetByEmail(inputUser.Email)
		if err == nil && userWithSameEmail != nil && userWithSameEmail.ID != inputUser.ID {
			return nil, errors.New("email already in use by another user")
		}
		existingUser.Email = inputUser.Email
	}

	// 5. Check if password changed
	if inputUser.PasswordHash != "" && inputUser.PasswordHash != existingUser.PasswordHash {
		// Validate password
		if err := s.policy.Validate(inputUser.PasswordHash); err != nil {
			return nil, errors.New("password is too weak: " + err.Error())
		}
		// Hash password
		hashedPassword, err := s.hasher.Hash(inputUser.PasswordHash)
		if err != nil {
			return nil, errors.New("failed to hash password: " + err.Error())
		}
		existingUser.PasswordHash = hashedPassword
	}

	// 6. Update time
	existingUser.UpdatedAt = time.Now()

	// 7. Update to DB
	if err := s.userRepo.Update(existingUser); err != nil {
		return nil, errors.New("failed to update user: " + err.Error())
	}

	// 8. Return mapped response
	return mapper.ToUserResponse(existingUser), nil
}

func (s *userService) DeleteUser(id int) error {
	// user id
	if id <= 0 {
		return errors.New("user id under zero")  
	}

	// get user form database
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return errors.New("failed to get user by ID: " + err.Error())
	}

	if user == nil {
		return errors.New("user not found")
	}

	// delete user from database
	if err := s.userRepo.Delete(id); err != nil {
		return errors.New("failed to delete user: " + err.Error())
	}

	// return nil if user deleted successfully
	return nil
}