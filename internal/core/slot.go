package core

import "parkingApp/pkg/models"

type Slot struct {
	number      int
	car         *models.Car
	isAvailable bool
}

func NewSlot(number int) *Slot {
	return &Slot{
		number:      number,
		isAvailable: true,
	}
}

func (s *Slot) Park(regNo string) {
	s.car = models.NewCar(regNo)
	s.isAvailable = false
}

func (s *Slot) Leave() {
	s.car = nil
	s.isAvailable = true
}

func (s *Slot) IsAvailable() bool {
	return s.isAvailable
}

func (s *Slot) GetCar() *models.Car {
	return s.car
}

func (s *Slot) GetNumber() int {
	return s.number
}
