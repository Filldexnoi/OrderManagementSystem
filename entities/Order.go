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
	order := &Order{TransactionId: o.TransactionId, OrderId: o.OrderId, Status: o.Status}
	if o.Status == "" {
		return nil, errors.New("invalid o status")
	}
	if o.Status == status {
		return order, nil
	}
	switch o.Status {

	case "New":
		if status == "Paid" {
			order.Status = status
			return order, nil
		}

	case "Paid":
		if status == "Processing" {
			order.Status = status
			return order, nil
		}
		if status == "Done" {
			order.Status = status
			return order, nil
		}

	case "Processing":
		if status == "Done" {
			order.Status = status
			return order, nil
		}
	}

	return order, fmt.Errorf("%w: from %s to %s", errors.New("invalid o status"), o.Status, status)
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
