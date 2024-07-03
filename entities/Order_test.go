package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangeStatus(t *testing.T) {
	tests := []struct {
		initialStatus  string
		newStatus      string
		expectedStatus string
		expectError    bool
	}{
		{"", "Paid", "", true},
		{"New", "Paid", "Paid", false},
		{"New", "Processing", "New", true},
		{"Paid", "Processing", "Processing", false},
		{"Paid", "Done", "Done", false},
		{"Processing", "Done", "Done", false},
		{"Done", "Paid", "Done", true},
	}

	for _, test := range tests {
		order := &Order{Status: test.initialStatus}
		updatedOrder, err := order.ChangeStatus(test.newStatus)
		if test.expectError {
			assert.Error(t, err, "expected an error but got none")
		} else {
			assert.NoError(t, err, "did not expect an error but got one")
		}
		assert.Equal(t, test.expectedStatus, updatedOrder.Status, "expected status does not match")
	}
}

func TestInitStatus(t *testing.T) {
	tests := []struct {
		initialStatus  string
		expectedStatus string
		expectError    bool
	}{
		{"", "New", false},
		{"New", "New", false},
		{"Done", "Done", false},
		{"Paid", "Paid", true},
	}

	for _, test := range tests {
		order := &Order{Status: test.initialStatus}
		updatedOrder, err := order.InitStatus()
		if test.expectError {
			assert.Error(t, err, "expected an error but got none")
		} else {
			assert.NoError(t, err, "did not expect an error but got one")
		}
		assert.Equal(t, test.expectedStatus, updatedOrder.Status, "expected status does not match")
	}
}
