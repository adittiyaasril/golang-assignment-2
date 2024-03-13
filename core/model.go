package core

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `gorm:"foreignKey:OrderID" json:"items"`
}

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Code        string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"-"`
}
