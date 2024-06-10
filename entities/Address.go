package entities

import (
	"time"
)

type Address struct {
	AddressID        uint
	AddressDetail    string
	AddressCreatedAt time.Time
	AddressUpdatedAt time.Time
}
