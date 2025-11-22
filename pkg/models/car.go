package models

type Car struct {
	RegistrationNumber string
	Color              string
}

func NewCar(regNo string) *Car {
	return &Car{
		RegistrationNumber: regNo,
	}
}

func (c *Car) GetRegistrationNumber() string {
	return c.RegistrationNumber
}
