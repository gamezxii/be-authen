package dto

type RequestOTPRequest struct {
	OperatorID string `json:"-" validate:"required,uuid4"`
	Username   string `json:"username" binding:"required" validate:"required,min=10,max=10,startwithprefix"`
}

type VerifyOTPRequest struct {
	OperatorID  string `json:"-" validate:"required,uuid4"`
	Username    string `json:"username" binding:"required" validate:"required,e164"`
	RefCode     string `json:"ref_code" binding:"required" validate:"required,alphanum,len=6"`
	ConfirmCode string `json:"confirm_code" binding:"required" validate:"required,numeric,len=6"`
}
