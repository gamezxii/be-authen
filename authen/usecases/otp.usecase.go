package usecases

import (
	"be-authen/authen/dto"
	"be-authen/authen/models"
	"be-authen/authen/repositories"
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"gorm.io/gorm"
)

type OTPUsecase interface {
	GetOTPRequests(page int, limit int) ([]models.OTPRequest, int64, error)
	GetOTPConfirms(page int, limit int) ([]models.OTPConfirm, int64, error)
	RequestOTP(ctx context.Context, req dto.RequestOTPRequest) (string, error)
	VerifyOTP(ctx context.Context, req dto.VerifyOTPRequest) (bool, error)
}

type otpUsecaseImpl struct {
	baseRepo  *repositories.BaseRepository
	otpRepo   repositories.OTPRepository
	snsClient *sns.SNS
}

func NewOTPUsecase(baseRepo *repositories.BaseRepository, otpRepo repositories.OTPRepository, snsClient *sns.SNS) OTPUsecase {
	return &otpUsecaseImpl{baseRepo: baseRepo, otpRepo: otpRepo, snsClient: snsClient}
}

func (u *otpUsecaseImpl) RequestOTP(ctx context.Context, req dto.RequestOTPRequest) (string, error) {
	refCode := generateRefCode(6)

	otpCode := fmt.Sprintf("%06d", rand.Intn(1000000))

	otpRequest := &models.OTPRequest{
		OperatorID:  req.OperatorID,
		Username:    req.Username,
		RefCode:     refCode,
		OtpCode:     otpCode,
		IsVerifyOTP: false,
	}
	if err := u.otpRepo.SaveOTPRequest(ctx, otpRequest); err != nil {
		return "", err
	}

	e164PhoneNumber := ConvertToE164(req.Username)

	message := fmt.Sprintf("<#> OTP %s (Ref:%s) ห้ามแจ้งรหัสกับบุคคลอื่นทุกกรณี", otpCode, refCode)

	_, err := u.snsClient.Publish(&sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(e164PhoneNumber),
	})
	if err != nil {
		return "", err
	}

	return refCode, nil
}

func (u *otpUsecaseImpl) VerifyOTP(ctx context.Context, req dto.VerifyOTPRequest) (bool, error) {
	otpRequest, err := u.otpRepo.GetOTPRequestByRefCode(ctx, req.RefCode, req.Username, req.OperatorID)
	if err != nil {
		return false, err
	}

	if !otpRequest.IsVerifyOTP && otpRequest.OtpCode == req.ConfirmCode {

		fmt.Printf("OTP Request: %+v\n", otpRequest)
		// หรือใช้ log.Printf
		log.Printf("OTP Request: %+v\n", otpRequest)

		err := u.baseRepo.WithTransaction(ctx, func(tx *gorm.DB) error {
			if err := u.otpRepo.UpdateOTPRequestStatus(ctx, tx, otpRequest.ID); err != nil {
				return err
			}

			otpConfirm := &models.OTPConfirm{
				OperatorID:  req.OperatorID,
				Username:    req.Username,
				ConfirmCode: otpRequest.OtpCode,
				IsUsed:      true,
			}

			if err := u.otpRepo.SaveOTPConfirm(ctx, tx, otpConfirm); err != nil {
				return err
			}

			return nil
		})

		// ตรวจสอบว่ามีข้อผิดพลาดจาก transaction หรือไม่
		if err != nil {
			return false, err
		}

		return true, nil // การตรวจสอบ OTP สำเร็จ
	}

	return false, nil // OTP ไม่ถูกต้องหรือถูกใช้ไปแล้ว
}

func (u *otpUsecaseImpl) GetOTPRequests(page int, limit int) ([]models.OTPRequest, int64, error) {
	return u.otpRepo.GetOTPRequests(page, limit)
}

func (u *otpUsecaseImpl) GetOTPConfirms(page int, limit int) ([]models.OTPConfirm, int64, error) {
	return u.otpRepo.GetOTPConfirms(page, limit)
}

func generateRefCode(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func ConvertToE164(phoneNumber string) string {
	if len(phoneNumber) == 10 && phoneNumber[0] == '0' {
		return "+66" + phoneNumber[1:]
	}
	return phoneNumber
}
