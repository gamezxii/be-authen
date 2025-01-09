package dto

type CreateUserRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	FullName   string `json:"full_name" binding:"required"`
	OperatorID string `json:"-"`
}
