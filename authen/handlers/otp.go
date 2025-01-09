package handlers

import (
	"be-authen/authen/commons"
	"be-authen/authen/dto"
	"be-authen/authen/usecases"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OTPHandler struct {
	otpUsecase usecases.OTPUsecase
	validate   *validator.Validate
}

func NewOTPHandler(otpUsecase usecases.OTPUsecase) *OTPHandler {
	validate := validator.New()

	// ลงทะเบียน custom validation สำหรับหมายเลขโทรศัพท์ที่เริ่มต้นด้วย 06, 08 หรือ 09
	validate.RegisterValidation("startwithprefix", func(fl validator.FieldLevel) bool {
		phone := fl.Field().String()
		// ตรวจสอบว่าหมายเลขโทรศัพท์เริ่มต้นด้วย 06, 08 หรือ 09
		re := regexp.MustCompile(`^(06|08|09)`)
		return re.MatchString(phone)
	})

	return &OTPHandler{
		otpUsecase: otpUsecase,
		validate:   validate,
	}
}

func (h *OTPHandler) RequestOTP(c *gin.Context) {
	var req dto.RequestOTPRequest
	req.OperatorID = commons.GetOperatorIDFromHeader(c)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.validate.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refCode, err := h.otpUsecase.RequestOTP(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to request OTP",
			"details": err.Error(), // แสดงรายละเอียดของ error
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ref_code": refCode})
}

func (h *OTPHandler) VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPRequest
	req.OperatorID = commons.GetOperatorIDFromHeader(c)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	isValid, err := h.otpUsecase.VerifyOTP(c, req)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify OTP"})
	// 	return
	// }
	if err != nil {
		// ตรวจสอบว่าเป็นข้อผิดพลาดที่ไม่พบข้อมูลหรือไม่
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "OTP request not found"})
		} else {
			// ข้อผิดพลาดอื่น ๆ ถือว่าเป็น internal server error
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to verify OTP",
				"details": err.Error(),
			})
		}
		return
	}

	if isValid {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid OTP"})
	}
}

func (h *OTPHandler) GetOtpRequests(c *gin.Context) {

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

	otpRequests, count, err := h.otpUsecase.GetOTPRequests(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  otpRequests,
		"count": count,
		"page":  page,
		"limit": limit,
	})
}

func (h *OTPHandler) GetOtpConfirms(c *gin.Context) {

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

	otpConfirms, count, err := h.otpUsecase.GetOTPConfirms(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  otpConfirms,
		"count": count,
		"page":  page,
		"limit": limit,
	})
}
