package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id           int       `json:"-" gorm:"primaryKey"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName" gorm:"type:varchar(191);not null"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Item struct {
	Id          int    `json:"-" gorm:"primaryKey"`
	ItemCode    string `json:"itemCode" gorm:"type:varchar(191);not null"`
	Description string `json:"description" gorm:"type:varchar(191);not null"`
	Quantity    int    `json:"quantity" gorm:"type:int;not null"`
	OrderId     int    `json:"-"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	if i.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	return
}
