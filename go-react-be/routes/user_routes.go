package routes

import (
	"bhachauk.github.io/go-react-sample/go-react-be/dao"
	"bhachauk.github.io/go-react-sample/go-react-be/dto"
	"bhachauk.github.io/go-react-sample/go-react-be/models"
	validation "bhachauk.github.io/go-react-sample/go-react-be/validaiton"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

// UserRoutes sets up the routes for user management
// It now takes the UserDAO directly
func UserRoutes(router *gin.Engine, userDAO *dao.UserDAO) {
	v1 := router.Group("/api/v1")
	{
		usersGroup := v1.Group("/users")
		{
			usersGroup.POST("/", func(c *gin.Context) {
				var req dto.CreateUserRequest
				if err := c.ShouldBindJSON(&req); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request body: %v", err.Error())})
					return
				}

				if err := validation.ValidateCreateUserRequest(&req, userDAO); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				user := models.User{
					Username: req.Username,
					Email:    req.Email,
					Password: req.Password, // In a real app, hash this!
				}

				if err := userDAO.CreateUser(&user); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
					return
				}

				c.JSON(http.StatusCreated, toUserResponse(&user))
			})

			usersGroup.GET("/", func(c *gin.Context) {
				users, err := userDAO.GetAllUsers()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
					return
				}
				var userResponses []dto.UserResponse
				for _, user := range users {
					userResponses = append(userResponses, toUserResponse(&user))
				}

				c.JSON(http.StatusOK, userResponses)
			})

			// Get User By ID
			// @Summary Get user by ID
			// @Description Retrieves a single user by their unique ID
			// @Tags Users
			// @Produce json
			// @Param id path int true "User ID"
			// @Success 200 {object} dto.UserResponse
			// @Failure 400 {object} gin.H "Invalid user ID"
			// @Failure 404 {object} gin.H "User not found"
			// @Failure 500 {object} gin.H "Internal server error"
			// @Router /users/{id} [get]
			usersGroup.GET("/:id", func(c *gin.Context) {
				idParam := c.Param("id")
				id, err := strconv.ParseUint(idParam, 10, 32)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
					return
				}

				user, err := userDAO.GetUserByID(uint(id))
				if err != nil {
					if err == gorm.ErrRecordNotFound {
						c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
						return
					}
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
					return
				}
				if user == nil { // Check if user was actually found (GORM returns nil for ErrRecordNotFound sometimes)
					c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
					return
				}

				c.JSON(http.StatusOK, toUserResponse(user))
			})

			// Update User
			// @Summary Update an existing user
			// @Description Updates an existing user's details by ID
			// @Tags Users
			// @Accept json
			// @Produce json
			// @Param id path int true "User ID"
			// @Param user body dto.UpdateUserRequest true "User update request"
			// @Success 200 {object} dto.UserResponse
			// @Failure 400 {object} gin.H "Invalid input"
			// @Failure 404 {object} gin.H "User not found"
			// @Failure 500 {object} gin.H "Internal server error"
			// @Router /users/{id} [put]
			usersGroup.PUT("/:id", func(c *gin.Context) {
				idParam := c.Param("id")
				id, err := strconv.ParseUint(idParam, 10, 32)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
					return
				}

				var req dto.UpdateUserRequest
				if err := c.ShouldBindJSON(&req); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request body: %v", err.Error())})
					return
				}

				existingUser, err := userDAO.GetUserByID(uint(id))
				if err != nil {
					if err == gorm.ErrRecordNotFound {
						c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
						return
					}
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
					return
				}
				if existingUser == nil {
					c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
					return
				}

				if err := validation.ValidateUpdateUserRequest(&req, existingUser.ID, userDAO); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if req.Username != "" {
					existingUser.Username = req.Username
				}
				if req.Email != "" {
					existingUser.Email = req.Email
				}

				if err := userDAO.UpdateUser(existingUser); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
					return
				}

				c.JSON(http.StatusOK, toUserResponse(existingUser))
			})

			// Delete User
			// @Summary Delete a user
			// @Description Deletes a user by their unique ID
			// @Tags Users
			// @Produce json
			// @Param id path int true "User ID"
			// @Success 204 "No Content"
			// @Failure 400 {object} gin.H "Invalid user ID"
			// @Failure 404 {object} gin.H "User not found"
			// @Failure 500 {object} gin.H "Internal server error"
			// @Router /users/{id} [delete]
			usersGroup.DELETE("/:id", func(c *gin.Context) {
				idParam := c.Param("id")
				id, err := strconv.ParseUint(idParam, 10, 32)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
					return
				}

				user, err := userDAO.GetUserByID(uint(id))
				if err != nil {
					if err == gorm.ErrRecordNotFound {
						c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
						return
					}
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user existence"})
					return
				}
				if user == nil {
					c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
					return
				}

				if err := userDAO.DeleteUser(uint(id)); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
					return
				}

				c.Status(http.StatusNoContent)
			})
		}
	}
}

// toUserResponse is a helper function, moved here since handlers are gone.
func toUserResponse(user *models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}
