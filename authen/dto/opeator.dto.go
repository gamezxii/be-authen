package dto

type CreateOperatorRequest struct {
	Name            string `json:"name" binding:"required" validate:"required"`
	IsActive        *bool  `json:"is_active" validate:"required"`
	IsMaintenance   *bool  `json:"is_maintenance" validate:"required"`
	IsVerifyOTP     *bool  `json:"is_verify_otp" validate:"required"`
	IsVerifyBank    *bool  `json:"is_verify_bank" validate:"required"`
	IsAllowWithdraw *bool  `json:"is_allow_withdraw" validate:"required"`
}

type UpdateOperatorRequest struct {
	Name            *string `json:"name" validate:"required,min=3"`
	IsActive        *bool   `json:"is_active" validate:"required"`
	IsMaintenance   *bool   `json:"is_maintenance" validate:"required"`
	IsVerifyOTP     *bool   `json:"is_verify_otp" validate:"required"`
	IsVerifyBank    *bool   `json:"is_verify_bank" validate:"required"`
	IsAllowWithdraw *bool   `json:"is_allow_withdraw" validate:"required"`
}
