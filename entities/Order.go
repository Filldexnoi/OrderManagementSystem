package entities

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Order struct {
	OrderId       uuid.UUID
	TransactionId uuid.UUID
	Status        string
}

func (o *Order) ChangeStatus(status string) (*Order, error) {
	if o.Status == "" {
		return o, errors.New("invalid o status")
	}
	if o.Status == status {
		return o, nil
	}
	switch o.Status {

	case "New":
		if status == "Paid" {
			o.Status = status
			return o, nil
		}

	case "Paid":
		if status == "Processing" {
			o.Status = status
			return o, nil
		}
		if status == "Done" {
			o.Status = status
			return o, nil
		}

	case "Processing":
		if status == "Done" {
			o.Status = status
			return o, nil
		}
	}

	return o, fmt.Errorf("%w: from %s to %s", errors.New("invalid o status"), o.Status, status)
}

func (o *Order) InitStatus() (*Order, error) {
	switch o.Status {
	case "":
		o.Status = "New"
		return o, nil
	case "New", "Done":
		return o, nil
	}

	return o, fmt.Errorf("%w: from %s to %s", errors.New("invalid order status"), o.Status, "New")
}

func (o *Order) IsEqualCreatedOrder(createdOrder *Order) bool {
	if o.TransactionId != createdOrder.TransactionId {
		return false
	}
	return true
}

func (o *Order) IsEqualUpdatedOrder(updateOrder *Order) bool {
	if o.OrderId != updateOrder.OrderId || o.Status != updateOrder.Status {
		return false
	}
	return true
}
