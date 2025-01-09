package repositories

import (
	"be-authen/authen/models"

	"gorm.io/gorm"
)

type OperatorRepository interface {
	GetAllOperatos(page int, limit int) ([]models.Operator, int64, error)
	CreateOperator(operator *models.Operator) error
	GetOperatorByID(id string) (*models.Operator, error)
	UpdateOperator(operator *models.Operator) error
	DeleteOperator(id string) error
}

type operatorRepositoryImpl struct {
	Db *gorm.DB
}

func NewOperatorRepository(db *gorm.DB) OperatorRepository {
	return &operatorRepositoryImpl{Db: db}
}

func (r *operatorRepositoryImpl) GetAllOperatos(page int, limit int) ([]models.Operator, int64, error) {
	var operators []models.Operator
	var count int64

	if err := r.Db.Model(&models.Operator{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	err := r.Db.Limit(limit).Offset(offset).Find(&operators).Error
	return operators, count, err
}

func (r *operatorRepositoryImpl) CreateOperator(operator *models.Operator) error {
	return r.Db.Create(operator).Error
}

func (r *operatorRepositoryImpl) GetOperatorByID(id string) (*models.Operator, error) {
	var operator models.Operator
	err := r.Db.First(&operator, "id = ?", id).Error
	return &operator, err
}

func (r *operatorRepositoryImpl) UpdateOperator(operator *models.Operator) error {
	return r.Db.Model(&models.Operator{}).
		Where("id = ?", operator.ID).
		Updates(operator).Error
}

func (r *operatorRepositoryImpl) DeleteOperator(id string) error {
	return r.Db.Where("id = ?", id).Delete(&models.Operator{}).Error
}
