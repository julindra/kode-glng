package models

import "time"

type Order struct {
	Order_id      uint      `gorm:"primaryKey" json:"orderId"`
	Customer_name string    `json:"customerName"`
	Ordered_at    time.Time `json:"orderedAt"`
	Items         []Item    `gorm:"foreignKey:order_id;references:order_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
