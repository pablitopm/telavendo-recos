package dynamodb

import (
	"sync"

	"telavendo.com.ar/telavendo-recos/app/domain/model"
)

type userPurchaseRepository struct {
	mu *sync.Mutex
	//games map[int]*model.Game
}

func NewUserPurchaseRepository() *userPurchaseRepository {
	return &userPurchaseRepository{
		mu: &sync.Mutex{},
		//games: map[int]*model.Game{},
	}
}

func (r *userPurchaseRepository) Upsert(*model.UserPurchase) error {
	return nil
}
