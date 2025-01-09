package repositories

import (
	"be-authen/authen/models"
	"context"

	"gorm.io/gorm"
)

type OTPRepository interface {
	SaveOTPRequest(ctx context.Context, otpRequest *models.OTPRequest) error
	UpdateOTPRequestStatus(ctx context.Context, tx *gorm.DB, id string) error
	GetOTPRequestByRefCode(ctx context.Context, refCode string, username string, operatorID string) (*models.OTPRequest, error)
	GetOTPRequests(page int, limit int) ([]models.OTPRequest, int64, error)
	SaveOTPConfirm(ctx context.Context, tx *gorm.DB, otpConfirm *models.OTPConfirm) error
	// UpdateOTPConfirmStatus(ctx context.Context, refCode string) error
	GetOTPConfirms(page int, limit int) ([]models.OTPConfirm, int64, error)
}

type otpRepositoryImpl struct {
	*BaseRepository
}

func NewOTPRepository(db *gorm.DB) OTPRepository {
	return &otpRepositoryImpl{
		BaseRepository: NewBaseRepository(db),
	}
}

// OTP REQUEST

func (r *otpRepositoryImpl) SaveOTPRequest(ctx context.Context, otpRequest *models.OTPRequest) error {
	return r.Db.WithContext(ctx).Create(otpRequest).Error
}

func (r *otpRepositoryImpl) UpdateOTPRequestStatus(ctx context.Context, tx *gorm.DB, id string) error {
	return r.Db.WithContext(ctx).Model(&models.OTPRequest{}).Where("id = ?", id).Update("is_verify_otp", true).Error
}

func (r *otpRepositoryImpl) GetOTPRequestByRefCode(ctx context.Context, refCode string, username string, operatorID string) (*models.OTPRequest, error) {
	var otpRequest models.OTPRequest
	err := r.Db.WithContext(ctx).
		Where("ref_code = ? AND username = ? AND operator_id = ? AND is_verify_otp = ?", refCode, username, operatorID, false).
		Order("created_at DESC").
		First(&otpRequest).Error
	return &otpRequest, err
}

func (r *otpRepositoryImpl) GetOTPRequests(page int, limit int) ([]models.OTPRequest, int64, error) {
	var otpRequests []models.OTPRequest
	var count int64

	if err := r.Db.Model(&models.OTPRequest{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	err := r.Db.Limit(limit).Offset(offset).Find(&otpRequests).Error
	return otpRequests, count, err
}

// OTP CONFIRM

func (r *otpRepositoryImpl) SaveOTPConfirm(ctx context.Context, tx *gorm.DB, otpConfirm *models.OTPConfirm) error {
	return r.Db.WithContext(ctx).Create(otpConfirm).Error
}

func (r *otpRepositoryImpl) GetOTPConfirms(page int, limit int) ([]models.OTPConfirm, int64, error) {
	var otpConfirms []models.OTPConfirm
	var count int64

	if err := r.Db.Model(&models.OTPConfirm{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	err := r.Db.Limit(limit).Offset(offset).Find(&otpConfirms).Error
	return otpConfirms, count, err
}

func (r *otpRepositoryImpl) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.Db.WithContext(ctx).Transaction(fn)
}
