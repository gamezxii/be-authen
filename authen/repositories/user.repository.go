package repositories

import (
	"be-authen/authen/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(page int, limit int) ([]models.User, int64, error)
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	SoftDeleteUser(id string, operatorID string) error
	SuspendUser(id string, operatorID string) error
}

type userRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{Db: db}
}

func (r *userRepositoryImpl) GetAllUsers(page int, limit int) ([]models.User, int64, error) {
	var users []models.User
	var count int64

	if err := r.Db.Model(&models.User{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	err := r.Db.Limit(limit).Offset(offset).Find(&users).Error
	return users, count, err
}

func (r *userRepositoryImpl) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.Db.First(&user, "id = ?", id).Error
	return &user, err
}

func (r *userRepositoryImpl) CreateUser(user *models.User) error {
	if err := r.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryImpl) UpdateUser(user *models.User) error {
	return r.Db.Save(user).Error
}

func (r *userRepositoryImpl) SuspendUser(id string, operatorID string) error {
	return r.Db.Model(&models.User{}).
		Where("id = ? AND operator_id = ?", id, operatorID).
		Update("is_active", false).Error
}

func (r *userRepositoryImpl) SoftDeleteUser(id string, operatorID string) error {
	return r.Db.Where("id = ? AND operator_id = ?", id, operatorID).Delete(&models.User{}).Error
}
