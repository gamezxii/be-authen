package usecases

import (
	"be-authen/authen/dto"
	"be-authen/authen/models"
	"be-authen/authen/repositories"
	"fmt"
)

type UserUsecase interface {
	GetAllUsers(page int, limit int) ([]models.User, int64, error)
	GetUserByID(id string) (*models.User, error)
	CreateUser(req *dto.CreateUserRequest) error
	// UpdateUser(id string, req *dto.UpdateUserRequest) error
	SoftDeleteUser(id string, operatorID string) error
	SuspendUser(id string, operatorID string) error
}

type userUsecaseImpl struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &userUsecaseImpl{
		userRepo: userRepo,
	}
}

func (u *userUsecaseImpl) GetAllUsers(page int, limit int) ([]models.User, int64, error) {
	return u.userRepo.GetAllUsers(page, limit)
}

// GetUserByID ดึงข้อมูลผู้ใช้ตาม ID
func (u *userUsecaseImpl) GetUserByID(id string) (*models.User, error) {
	return u.userRepo.GetUserByID(id)
}

// UpdateUser อัปเดตข้อมูลผู้ใช้
// func (u *userUsecaseImpl) UpdateUser(id string, req *dto.UpdateUserRequest) error {
// 	user, err := u.userRepo.GetUserByID(id)
// 	if err != nil {
// 		return err
// 	}

// 	user.FullName = req.FullName
// 	user.IsActive = req.IsActive

// 	return u.userRepo.UpdateUser(user)
// }

func (u *userUsecaseImpl) CreateUser(req *dto.CreateUserRequest) error {
	// Business Logic
	if req.Username == "" || req.Password == "" {
		return fmt.Errorf("username and password are required")
	}

	if req.FullName == "" {
		return fmt.Errorf("FullName must be not empty")
	}

	newUser := models.User{
		Username:   req.Username,
		Password:   req.Password,
		FullName:   req.FullName,
		OperatorID: req.OperatorID,
	}

	return u.userRepo.CreateUser(&newUser)
}

func (u *userUsecaseImpl) SuspendUser(id string, operatorID string) error {
	return u.userRepo.SuspendUser(id, operatorID)
}

func (u *userUsecaseImpl) SoftDeleteUser(id string, operatorID string) error {
	return u.userRepo.SoftDeleteUser(id, operatorID)
}
