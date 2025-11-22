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
		return ""
	}

	switch cmd.Name {
	case "create_parking_lot":
		capacity, _ := strconv.Atoi(cmd.Args[0])
		return h.parkingService.CreateParkingLot(capacity)

	case "park":
		req := &dto.ParkRequest{RegistrationNumber: cmd.Args[0]}
		return h.parkingService.ParkCar(req)

	case "leave":
		hours, _ := strconv.Atoi(cmd.Args[1])
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
