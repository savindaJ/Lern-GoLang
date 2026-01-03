package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	service UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error" example:"error message"`
	Details string `json:"details,omitempty" example:"additional details"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
	Message string `json:"message" example:"operation successful"`
}

// PaginatedResponse represents a paginated response
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total" example:"100"`
	Page       int         `json:"page" example:"1"`
	Limit      int         `json:"limit" example:"10"`
	TotalPages int         `json:"total_pages" example:"10"`
}

// Register godoc
// @Summary      Register a new user
// @Description  Create a new user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body CreateUserRequest true "User registration data"
// @Success      201  {object}  UserResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      409  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /users/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}

	user, err := h.service.Register(&req)
	if err != nil {
		if err == ErrEmailAlreadyExists {
			c.JSON(http.StatusConflict, ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Login godoc
// @Summary      User login
// @Description  Authenticate a user and return a token
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "User login credentials"
// @Success      200  {object}  LoginResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}

	user, err := h.service.Login(&req)
	if err != nil {
		if err == ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Login failed"})
		return
	}

	// TODO: Generate JWT token here
	response := LoginResponse{
		Token: "jwt-token-placeholder",
		User:  user.ToResponse(),
	}

	c.JSON(http.StatusOK, response)
}

// GetAll godoc
// @Summary      Get all users
// @Description  Retrieve all users with pagination
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        page   query     int  false  "Page number"  default(1)
// @Param        limit  query     int  false  "Items per page"  default(10)
// @Success      200  {object}  PaginatedResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /users [get]
func (h *UserHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	users, total, err := h.service.GetAll(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to fetch users"})
		return
	}

	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, PaginatedResponse{
		Data:       users,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	})
}

// GetByID godoc
// @Summary      Get user by ID
// @Description  Retrieve a user by their ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  UserResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		if err == ErrUserNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Update godoc
// @Summary      Update user
// @Description  Update an existing user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id      path      int                true  "User ID"
// @Param        request body      UpdateUserRequest  true  "User update data"
// @Success      200  {object}  UserResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      409  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request", Details: err.Error()})
		return
	}

	user, err := h.service.Update(uint(id), &req)
	if err != nil {
		if err == ErrUserNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
			return
		}
		if err == ErrEmailAlreadyExists {
			c.JSON(http.StatusConflict, ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Delete godoc
// @Summary      Delete user
// @Description  Delete a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  SuccessResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		if err == ErrUserNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "User deleted successfully"})
}
