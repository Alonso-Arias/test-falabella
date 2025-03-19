package model

type Book struct {
	ID        uint    `json:"id"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Publisher string  `json:"publisher"`
	Country   string  `json:"country"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
}

type GetBookResponse struct {
	TotalPrice int `json:"totalPrice"`
}
