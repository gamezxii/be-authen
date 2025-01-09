package di

import (
	"be-authen/authen/handlers"
	"be-authen/authen/repositories"
	"be-authen/authen/usecases"

	"github.com/aws/aws-sdk-go/service/sns"
	"gorm.io/gorm"
)

type Container struct {
	UserHandler     *handlers.UserHandler
	OperatorHandler *handlers.OperatorHandler
	OtpHandler      *handlers.OTPHandler
}

func NewContainer(db *gorm.DB, snsClient *sns.SNS) *Container {

	baseRepo := repositories.NewBaseRepository(db)

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userHandler := handlers.NewUserHandler(userUsecase)

	operatorRepository := repositories.NewOperatorRepository(db)
	operatorUsecase := usecases.NewOperatorUsecase(operatorRepository)
	operatorHandler := handlers.NewOperatorHandler(operatorUsecase)

	otpRepository := repositories.NewOTPRepository(db)
	otpUsecase := usecases.NewOTPUsecase(baseRepo, otpRepository, snsClient)
	otpHandler := handlers.NewOTPHandler(otpUsecase)

	return &Container{
		UserHandler:     userHandler,
		OperatorHandler: operatorHandler,
		OtpHandler:      otpHandler,
	}
}
