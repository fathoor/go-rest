package models

import "time"

type Order struct {
	Id           int       `json:"orderId" gorm:"primaryKey"`
	CustomerName string    `json:"customerName" gorm:"type:varchar(191);not null"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderId"`
}

type Item struct {
	Id          int    `json:"itemId" gorm:"primaryKey"`
	ItemCode    string `json:"itemCode" gorm:"type:varchar(191);not null"`
	Description string `json:"description" gorm:"type:varchar(191);not null"`
	Quantity    int    `json:"quantity" gorm:"type:int;not null"`
	OrderId     int
}
