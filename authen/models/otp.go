package models

import (
	"time"

	"gorm.io/gorm"
)

type OTPRequest struct {
	ID          string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OperatorID  string         `gorm:"type:varchar(255);not null"`
	Username    string         `gorm:"type:varchar(255);not null"`
	RefCode     string         `gorm:"type:varchar(255);not null"`
	OtpCode     string         `gorm:"type:varchar(255);not null"`
	IsVerifyOTP bool           `gorm:"default:false"`
	CreatedAt   time.Time      `gorm:"type:timestamptz;default:now()"`
	UpdatedAt   time.Time      `gorm:"type:timestamptz;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (OTPRequest) TableName() string {
	return "otp_request"
}

type OTPConfirm struct {
	ID          string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OperatorID  string         `gorm:"type:varchar(255);not null"`
	Username    string         `gorm:"type:varchar(255);not null"`
	ConfirmCode string         `gorm:"type:varchar(255);not null"`
	IsUsed      bool           `gorm:"default:false"`
	CreatedAt   time.Time      `gorm:"type:timestamptz;default:now()"`
	UpdatedAt   time.Time      `gorm:"type:timestamptz;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (OTPConfirm) TableName() string {
	return "otp_confirm"
}
