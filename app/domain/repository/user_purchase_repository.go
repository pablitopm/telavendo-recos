package repository

import (
	"telavendo.com.ar/telavendo-recos/app/domain/model"
)

type UserPurchaseRepository interface {
	Upsert(purchase *model.UserPurchase) error
	Get(purchase *model.UserPurchase) (*model.UserPurchase, error)
}
