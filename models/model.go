package models

import "time"

type Order struct {
	Id           int       `json:"-" gorm:"primaryKey"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName" gorm:"type:varchar(191);not null"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderId"`
}

type Item struct {
	Id          int    `json:"-" gorm:"primaryKey"`
	ItemCode    string `json:"itemCode" gorm:"type:varchar(191);not null"`
	Description string `json:"description" gorm:"type:varchar(191);not null"`
	Quantity    int    `json:"quantity" gorm:"type:int;not null"`
	OrderId     int    `json:"-"`
}
