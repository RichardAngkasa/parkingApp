package dto

type ParkRequest struct {
	RegistrationNumber string
}

type LeaveRequest struct {
	RegistrationNumber string
	Hours              int
}
