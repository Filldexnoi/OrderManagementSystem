package entities

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalPrice(t *testing.T) {
	tests := []struct {
		name          string
		items         []Item
		address       string
		expectedPrice float64
	}{
		{
			name: "Within Thailand",
			items: []Item{
				{ProductId: 1, Quantity: 1, Price: 100.0},
				{ProductId: 2, Quantity: 2, Price: 50.0},
			},
			address:       "th",
			expectedPrice: 250.0,
		},
		{
			name: "Outside Thailand",
			items: []Item{
				{ProductId: 1, Quantity: 1, Price: 100.0},
				{ProductId: 2, Quantity: 2, Price: 50.0},
			},
			address:       "us",
			expectedPrice: 400.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transaction := &Transaction{
				TransactionId: uuid.New(),
				OrderAddress:  tt.address,
				Items:         tt.items,
			}
			transaction.CalPrice()
			assert.Equal(t, tt.expectedPrice, transaction.TotalPrice)
		})
	}
}

func TestIsWithinThailand(t *testing.T) {
	transaction := &Transaction{}

	assert.True(t, transaction.isWithinThailand("th"), "Expected true for Thailand")
	assert.False(t, transaction.isWithinThailand("us"), "Expected false for non-Thailand")
}

func TestIsValidCountry(t *testing.T) {
	transaction := &Transaction{}

	assert.True(t, transaction.IsValidCountry("th"), "Expected true for valid country code")
	assert.False(t, transaction.IsValidCountry("xx"), "Expected false for invalid country code")
}
