package repository

import (
	"telavendo.com.ar/telavendo-recos/app/domain/model"
)

type ProductCombinationRankingRepository interface {
	Upsert(*model.ProductCombinationRanking) error
}
