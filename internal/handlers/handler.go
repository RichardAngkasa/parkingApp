package handlers

import (
	"parkingApp/internal/services"
	"parkingApp/pkg/commands"
	"parkingApp/pkg/dto"
	"strconv"
)

type Handler struct {
	parkingService *services.ParkingService
}

func NewHandler() *Handler {
	return &Handler{
		parkingService: services.NewParkingService(),
	}
}

func (h *Handler) HandleCommand(input string) string {
	cmd := commands.ParseInput(input)
	if cmd == nil {
		return "Error: Invalid command format"
	}

	switch cmd.Name {
	case "create_parking_lot":
		if len(cmd.Args) < 1 {
			return "Error: Missing capacity parameter"
		}
		capacity, err := strconv.Atoi(cmd.Args[0])
		if err != nil || capacity <= 0 {
			return "Error: Invalid capacity value"
		}
		return h.parkingService.CreateParkingLot(capacity)

	case "park":
		if len(cmd.Args) < 1 {
			return "Error: Missing registration number"
		}
		req := &dto.ParkRequest{RegistrationNumber: cmd.Args[0]}
		return h.parkingService.ParkCar(req)

	case "leave":
		if len(cmd.Args) < 2 {
			return "Error: Missing hours parameter"
		}
		hours, err := strconv.Atoi(cmd.Args[1])
		if err != nil || hours <= 0 {
			return "Error: Invalid hours value"
		}
		req := &dto.LeaveRequest{
			RegistrationNumber: cmd.Args[0],
			Hours:              hours,
		}
		return h.parkingService.LeaveCar(req)

	case "status":
		return h.parkingService.GetStatus()

	default:
		return h.parkingService.HandleInvalidCommand()
	}
}
