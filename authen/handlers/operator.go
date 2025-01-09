package handlers

import (
	"net/http"
	"strconv"

	"be-authen/authen/dto"
	"be-authen/authen/models"
	"be-authen/authen/usecases"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OperatorHandler struct {
	operatorUsecase usecases.OperatorUsecase
	validate        *validator.Validate
}

func NewOperatorHandler(operatorUsecase usecases.OperatorUsecase) *OperatorHandler {
	return &OperatorHandler{operatorUsecase: operatorUsecase}
}

func (h *OperatorHandler) GetAllOperatos(c *gin.Context) {

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

	operators, count, err := h.operatorUsecase.GetAllOperatos(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  operators,
		"count": count,
		"page":  page,
		"limit": limit,
	})
}

func (h *OperatorHandler) GetOperatorByID(c *gin.Context) {
	operatorID := c.Param("id")

	operator, err := h.operatorUsecase.GetOperatorByID(operatorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Operator not found"})
		return
	}

	c.JSON(http.StatusOK, operator)
}

func (h *OperatorHandler) CreateOperator(c *gin.Context) {
	var req dto.CreateOperatorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Map DTO to model
	operator := models.Operator{
		Name:            req.Name,
		IsActive:        *req.IsActive,
		IsMaintenance:   *req.IsMaintenance,
		IsVerifyOTP:     *req.IsVerifyOTP,
		IsVerifyBank:    *req.IsVerifyBank,
		IsAllowWithdraw: *req.IsAllowWithdraw,
	}

	if err := operator.GenerateKeys(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate keys"})
		return
	}

	createdOperator, err := h.operatorUsecase.CreateOperator(&operator)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create operator"})
		return
	}

	c.JSON(http.StatusCreated, createdOperator)

}

func (h *OperatorHandler) UpdateOperator(c *gin.Context) {
	var req dto.UpdateOperatorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// ดึง operator ที่ต้องการอัปเดต
	operatorID := c.Param("id")
	operator, err := h.operatorUsecase.GetOperatorByID(operatorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Operator not found"})
		return
	}

	// Map DTO ไปยังฟิลด์ใน model
	if req.Name != nil {
		operator.Name = *req.Name
	}
	if req.IsActive != nil {
		operator.IsActive = *req.IsActive
	}
	if req.IsMaintenance != nil {
		operator.IsMaintenance = *req.IsMaintenance
	}
	if req.IsVerifyOTP != nil {
		operator.IsVerifyOTP = *req.IsVerifyOTP
	}
	if req.IsVerifyBank != nil {
		operator.IsVerifyBank = *req.IsVerifyBank
	}
	if req.IsAllowWithdraw != nil {
		operator.IsAllowWithdraw = *req.IsAllowWithdraw
	}

	updatedOperator, err := h.operatorUsecase.UpdateOperator(operator)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update operator"})
		return
	}

	c.JSON(http.StatusCreated, updatedOperator)
}

func (h *OperatorHandler) DeleteOperator(c *gin.Context) {
	operatorID := c.Param("id")

	if err := h.operatorUsecase.DeleteOperator(operatorID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete operator"})
		return
	}

	c.Status(http.StatusNoContent)
}
