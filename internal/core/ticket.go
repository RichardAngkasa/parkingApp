package core

import "parkingApp/pkg/models"

type Ticket struct {
	slotNumber int
	car        *models.Car
}

func NewTicket(slotNumber int, car *models.Car) *Ticket {
	return &Ticket{
		slotNumber: slotNumber,
		car:        car,
	}
}

func CalculateCharge(hours int) int {
	if hours <= 2 {
		return 10
	}
	return 10 + (hours-2)*10
}
