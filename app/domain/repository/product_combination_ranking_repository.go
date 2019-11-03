package repository

import (
	"telavendo.com.ar/telavendo-recos/app/domain/model"
)

type ProductCombinationRepository interface {
	Upsert(combination *model.ProductCombination) error
	Get(combination *model.ProductCombination) (*model.ProductCombination, error)
}
