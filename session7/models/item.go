package models

type Item struct {
	Item_id     uint   `gorm:"primaryKey" json:"lineItemId"`
	Item_code   string `json:"itemCode"`
	Description string
	Quantity    int
	Order_id    uint `json:"orderId"`
}
