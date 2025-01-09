package models

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"time"

	"gorm.io/gorm"
)

type Operator struct {
	ID              string         `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name            string         `gorm:"type:varchar(255);not null"`
	IsActive        bool           `gorm:"default:true"`
	IsMaintenance   bool           `gorm:"default:false"`
	IsVerifyOTP     bool           `gorm:"default:false"`
	IsVerifyBank    bool           `gorm:"default:false"`
	IsAllowWithdraw bool           `gorm:"default:true"`
	PublicKey       string         `gorm:"type:text"`
	PrivateKey      string         `gorm:"type:text"`
	CreatedAt       time.Time      `gorm:"type:timestamptz;default:now()"`
	UpdatedAt       time.Time      `gorm:"type:timestamptz;default:now()"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (Operator) TableName() string {
	return "operator"
}

// GenerateKeys generates a new RSA public and private key pair
func (op *Operator) GenerateKeys() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// Encode private key to PEM format
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes})
	op.PrivateKey = string(privateKeyPEM)

	// Generate public key
	publicKey := &privateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: publicKeyBytes})
	op.PublicKey = string(publicKeyPEM)

	return nil
}
