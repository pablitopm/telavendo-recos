package repository

import (
	"telavendo.com.ar/telavendo-recos/app/domain/model"
)

type UserPurchaseRepository interface {
	Upsert(*model.UserPurchase) error
	//FindByUserID(userID int) ([]*model.UserPurchase, error)
}
