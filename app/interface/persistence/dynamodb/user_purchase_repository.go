package dynamodb

import (
	"fmt"
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

func (r *userPurchaseRepository) Upsert(purchase *model.UserPurchase) error {
	existingPurchase, err := r.Get(purchase)
	if err != nil {
		return err
	}
	if existingPurchase != nil {
		existingPurchase.PurchasedTimes++
	}

	fmt.Println("saved ")
	return nil
}

func (r *userPurchaseRepository) Get(purchase *model.UserPurchase) (*model.UserPurchase, error) {
	fmt.Println("get ")
	return nil, nil
}
