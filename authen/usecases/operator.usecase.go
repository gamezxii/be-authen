package usecases

import (
	"be-authen/authen/models"
	"be-authen/authen/repositories"
)

type OperatorUsecase interface {
	GetAllOperatos(page int, limit int) ([]models.Operator, int64, error)
	GetOperatorByID(id string) (*models.Operator, error)
	CreateOperator(operator *models.Operator) (*models.Operator, error)
	UpdateOperator(operator *models.Operator) (*models.Operator, error)
	DeleteOperator(id string) error
}

type operatorUsecaseImpl struct {
	operatorRepo repositories.OperatorRepository
}

func NewOperatorUsecase(operatorRepo repositories.OperatorRepository) OperatorUsecase {
	return &operatorUsecaseImpl{operatorRepo: operatorRepo}
}

func (u *operatorUsecaseImpl) GetAllOperatos(page int, limit int) ([]models.Operator, int64, error) {
	return u.operatorRepo.GetAllOperatos(page, limit)
}

func (u *operatorUsecaseImpl) CreateOperator(operator *models.Operator) (*models.Operator, error) {
	if err := u.operatorRepo.CreateOperator(operator); err != nil {
		return nil, err
	}
	return operator, nil
}

func (u *operatorUsecaseImpl) GetOperatorByID(id string) (*models.Operator, error) {
	return u.operatorRepo.GetOperatorByID(id)
}

func (u *operatorUsecaseImpl) UpdateOperator(operator *models.Operator) (*models.Operator, error) {
	if err := u.operatorRepo.UpdateOperator(operator); err != nil {
		return nil, err
	}
	return operator, nil
}

func (u *operatorUsecaseImpl) DeleteOperator(id string) error {
	return u.operatorRepo.DeleteOperator(id)
}
