package models

import (
	"time"
)

type Order struct {
	OrderId   uint        `gorm:"column:order_id;primary_key;AUTO_INCREMENT"`
	Address   string      `gorm:"column:order_address"`
	Items     []ItemModel `gorm:"foreignkey:order_id"`
	State     string      `gorm:"column:state"`
	CreatedAt time.Time   `gorm:"column:created_at"`
	UpdatedAt time.Time   `gorm:"column:updated_at"`
}
