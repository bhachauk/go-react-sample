package validation

import (
	"bhachauk.github.io/go-react-sample/go-react-be/dao"
	"bhachauk.github.io/go-react-sample/go-react-be/dto"
	"fmt"
	"github.com/go-playground/validator/v10"
)

// Validate struct for user DTOs
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateCreateUserRequest validates the CreateUserRequest DTO
func ValidateCreateUserRequest(req *dto.CreateUserRequest, userDAO *dao.UserDAO) error {
	if err := validate.Struct(req); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	// Custom validation for uniqueness
	if record, err := userDAO.GetUserByEmail(req.Email); err == nil {
		if record != nil {
			return fmt.Errorf("email '%s' is already registered", req.Email)
		}
	}
	if record, err := userDAO.GetUserByUsername(req.Username); err == nil {
		if record != nil {
			return fmt.Errorf("username '%s' is already taken", req.Username)
		}
	}
	return nil
}

// ValidateUpdateUserRequest validates the UpdateUserRequest DTO
func ValidateUpdateUserRequest(req *dto.UpdateUserRequest, currentUserID uint, userDAO *dao.UserDAO) error {
	if err := validate.Struct(req); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	// Custom validation for uniqueness during update
	if req.Email != "" {
		if user, err := userDAO.GetUserByEmail(req.Email); err == nil && user != nil && user.ID != currentUserID {
			return fmt.Errorf("email '%s' is already registered by another user", req.Email)
		}
	}
	if req.Username != "" {
		if user, err := userDAO.GetUserByUsername(req.Username); err == nil && user != nil && user.ID != currentUserID {
			return fmt.Errorf("username '%s' is already taken by another user", req.Username)
		}
	}
	return nil
}
