package models

import (
	"awesomeProject/entities"
	"github.com/google/uuid"
	"time"
)

type Order struct {
	OrderId       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TransactionID uuid.UUID `gorm:"type:uuid;not null;unique"`
	Status        string    `gorm:"column:status"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (*Order) TableName() string {
	return "orders"
}

func (o *Order) ToOrder() *entities.Order {
	return &entities.Order{
		OrderId:       o.OrderId,
		TransactionId: o.TransactionID,
		Status:        o.Status,
	}
}

func OrderToGormOrder(o entities.Order) Order {
	return Order{
		OrderId:       o.OrderId,
		TransactionID: o.TransactionId,
		Status:        o.Status,
	}
}
