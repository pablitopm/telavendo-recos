package model

type UserPurchase struct {
	UserID         int `json:"userId"`
	ProductID      int `json:"productId"`
	PurchasedTimes int `json:"purchasedTimes"`
}
