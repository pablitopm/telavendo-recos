package usecase

import (
	"telavendo.com.ar/telavendo-recos/app/domain/model"
	"telavendo.com.ar/telavendo-recos/app/domain/repository"
	"telavendo.com.ar/telavendo-recos/app/domain/service"
)

type UserPurchaseUsecase interface {
	RegisterUserPurchase(userID, productID int) error
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

func (u *userPurchaseUsecase) RegisterUserPurchase(userID, productID int) error {
	purchase := &model.UserPurchase{}
	purchase.ProductID = productID
	purchase.UserID = userID
	err := u.repo.Upsert(purchase)
	if err != nil {
		return err
	}
	return nil
}
