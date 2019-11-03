package model

//this are
type Order struct {
	Customer Customer  `json:"customer"`
	Products []Product `json:"products"`
}

type Product struct {
	ProductID int `json:"product_id"`
}

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
