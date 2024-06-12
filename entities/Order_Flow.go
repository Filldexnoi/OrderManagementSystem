package entities

import (
	"errors"
	"fmt"
)

type OrderFlow struct{}

func (f OrderFlow) ChangeStatus(order Order, status string) (Order, error) {
	if order.Status == "" {
		return order, errors.New("invalid order status")
	}

	if order.Status == status {
		return order, nil
	}

	switch order.Status {
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

	return order, fmt.Errorf("%w: from %s to %s", errors.New("invalid order status"), order.Status, status)
}

func (f OrderFlow) Init(order Order) (Order, error) {
	switch order.Status {
	case "":
		order.Status = "New"
		return order, nil
	case "New", "Done":
		return order, nil
	}

	return order, fmt.Errorf("%w: from %s to %s", errors.New("invalid order status"), order.Status, "New")
}
