package handlers

import (
	"net/http"
	"strconv"

	"be-authen/authen/commons"
	"be-authen/authen/dto"
	"be-authen/authen/usecases"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(userUsecase usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {

	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	users, count, err := h.userUsecase.GetAllUsers(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"count": count,
		"page":  page,
		"limit": limit,
	})
}

func (h *UserHandler) GetUserDetail(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.userUsecase.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	req.OperatorID = commons.GetOperatorIDFromHeader(c)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.userUsecase.CreateUser(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *UserHandler) SoftDeleteUser(c *gin.Context) {
	userID := c.Param("id")
	operatorID := commons.GetOperatorIDFromHeader(c)

	if err := h.userUsecase.SoftDeleteUser(userID, operatorID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *UserHandler) SuspendUser(c *gin.Context) {
	userID := c.Param("id")
	operatorID := commons.GetOperatorIDFromHeader(c)

	if err := h.userUsecase.SuspendUser(userID, operatorID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to suspend user"})
		return
	}

	c.Status(http.StatusNoContent)
}
