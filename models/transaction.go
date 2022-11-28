package models

type Transaction struct {
	ID        int             `json:"id" gorm:"primary_key:auto_increment"`
	ProductID int             `json:"product_id"`
	Product   ProductResponse `json:"product"`
	BuyerID   int             `json:"buyer_id"`
	Buyer     UsersResponse   `json:"buyer"`
	Price     int             `json:"price"`
	Status    string          `json:"status"  gorm:"type:varchar(25)"`
}
