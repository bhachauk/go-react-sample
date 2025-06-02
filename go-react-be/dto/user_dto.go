package dto

// CreateUserRequest defines the structure for creating a new user
type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=1,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

// UpdateUserRequest defines the structure for updating an existing user
type UpdateUserRequest struct {
	Username string `json:"username" validate:"omitempty,min=1,max=50"`
	Email    string `json:"email" validate:"omitempty,email"`
	// Password is not typically updated via a general update endpoint
	// A separate "change password" endpoint is usually preferred
}

// UserResponse defines the structure for sending user data to the client
type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
