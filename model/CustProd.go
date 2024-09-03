package model

type Customer struct {
	Name     string `json:"name"`
	Phn      int    `json:"phn"`
	Email    string `json:"email"`
	Quantity int    `json:"quantity"`
	Product  string `json:"product"`
}
type Product struct {
	ProdName string  `json:"Pname"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
