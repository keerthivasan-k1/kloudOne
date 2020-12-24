package model

type Product struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Price string `json:"price"`
}
