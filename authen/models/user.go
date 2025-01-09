package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         string         `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	OperatorID string         `gorm:"type:uuid;not null"`
	Username   string         `gorm:"type:varchar(255);not null"`
	Password   string         `gorm:"type:varchar(255);not null"`
	FullName   string         `gorm:"type:varchar(255)"`
	IsActive   bool           `gorm:"default:true"`
	LoginIP    string         `gorm:"type:varchar(255)"`
	CreatedAt  time.Time      `gorm:"type:timestamptz;default:now()"`
	UpdatedAt  time.Time      `gorm:"type:timestamptz;default:now()"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "user"
}
