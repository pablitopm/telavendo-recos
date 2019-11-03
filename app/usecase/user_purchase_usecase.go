package usecase

import (
	"telavendo.com.ar/telavendo-recos/app/domain/model"
	"telavendo.com.ar/telavendo-recos/app/domain/repository"
	"telavendo.com.ar/telavendo-recos/app/domain/service"
)

type UserPurchaseUsecase interface {
	RegisterUserPurchase(model.UserPurchase) error
}

type userPurchaseUsecase struct {
	repo    repository.UserPurchaseRepository
	service *service.UserPurchaseService
}

func NewUserPurchaseUsecase(repo repository.UserPurchaseRepository, service *service.UserPurchaseService) *userPurchaseUsecase {
	return &userPurchaseUsecase{
		repo:    repo,
		service: service,
	}
}

func (u *userPurchaseUsecase) RegisterUserPurchase(purchase model.UserPurchase) error {
	err := u.repo.Upsert(&purchase)
	if err != nil {
		return err
	}
	return nil
}
