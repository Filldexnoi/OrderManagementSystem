package models

import (
	"time"
)

type Address struct {
	AddressID        uint `gorm:"primaryKey"`
	AddressDetail    string
	AddressCreatedAt time.Time
	AddressUpdatedAt time.Time
}
