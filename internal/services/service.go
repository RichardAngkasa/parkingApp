package services

import (
	"parkingApp/internal/core"
	"parkingApp/internal/errors"
	"parkingApp/pkg/dto"
	"strconv"
)

type ParkingService struct {
	parkingLot *core.ParkingLot
}

func NewParkingService() *ParkingService {
	return &ParkingService{}
}

func (s *ParkingService) CreateParkingLot(capacity int) string {
	s.parkingLot = core.NewParkingLot(capacity)
	return "Created parking lot with " + strconv.Itoa(capacity) + " slots"
}

func (s *ParkingService) ParkCar(req *dto.ParkRequest) string {
	if s.parkingLot == nil {
		return errors.ErrParkingLotNotCreated
	}
	if s.isCarAlreadyParked(req.RegistrationNumber) {
		return "Error: Car " + req.RegistrationNumber + " is already parked"
	}
	if s.parkingLot.IsFull() {
		return errors.ErrParkingLotFull
	}

	slot := s.parkingLot.Park(req.RegistrationNumber)
	return "Allocated slot number: " + strconv.Itoa(slot)
}

func (s *ParkingService) isCarAlreadyParked(regNo string) bool {
	occupied := s.parkingLot.GetStatus()
	for _, slot := range occupied {
		if slot.GetCar().GetRegistrationNumber() == regNo {
			return true
		}
	}
	return false
}

func (s *ParkingService) LeaveCar(req *dto.LeaveRequest) string {
	if s.parkingLot == nil {
		return errors.ErrParkingLotNotCreated
	}

	slotNumber := s.parkingLot.Leave(req.RegistrationNumber)
	if slotNumber == -1 {
		return "Registration number " + req.RegistrationNumber + " not found"
	}

	charge := core.CalculateCharge(req.Hours)
	return "Registration number " + req.RegistrationNumber + " with Slot Number " + strconv.Itoa(slotNumber) + " is free with Charge $" + strconv.Itoa(charge)
}

func (s *ParkingService) GetStatus() string {
	if s.parkingLot == nil {
		return errors.ErrParkingLotNotCreated
	}

	slots := s.parkingLot.GetStatus()
	result := "Slot No.\tRegistration No.\n"
	for _, slot := range slots {
		result += strconv.Itoa(slot.GetNumber()) + "\t\t" + slot.GetCar().GetRegistrationNumber() + "\n"
	}
	return result
}

func (s *ParkingService) HandleInvalidCommand() string {
	return errors.ErrInvalidCommand
}
