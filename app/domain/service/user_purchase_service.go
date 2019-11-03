package service

import (
	"telavendo.com.ar/telavendo-recos/app/domain/repository"
)

type UserPurchaseService struct {
	repo repository.UserPurchaseRepository
}

func NewUserPurchaseService(repo repository.UserPurchaseRepository) *UserPurchaseService {
	return &UserPurchaseService{
		repo: repo,
	}
}
