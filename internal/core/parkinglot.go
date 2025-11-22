package core

type ParkingLot struct {
	capacity int
	slots    []*Slot
}

func NewParkingLot(capacity int) *ParkingLot {
	slots := make([]*Slot, capacity)
	for i := range slots {
		slots[i] = NewSlot(i + 1)
	}
	return &ParkingLot{
		capacity: capacity,
		slots:    slots,
	}
}

func (p *ParkingLot) Park(regNo string) int {
	for _, slot := range p.slots {
		if slot.IsAvailable() {
			slot.Park(regNo)
			ticket := NewTicket(slot.number, slot.car)
			return ticket.slotNumber
		}
	}
	return -1
}

func (p *ParkingLot) Leave(regNo string) int {
	for _, slot := range p.slots {
		if !slot.IsAvailable() && slot.car.RegistrationNumber == regNo {
			slot.Leave()
			return slot.number
		}
	}
	return -1
}

func (p *ParkingLot) GetStatus() []*Slot {
	var occupied []*Slot
	for _, slot := range p.slots {
		if !slot.IsAvailable() {
			occupied = append(occupied, slot)
		}
	}
	return occupied
}

func (p *ParkingLot) IsFull() bool {
	for _, slot := range p.slots {
		if slot.IsAvailable() {
			return false
		}
	}
	return true
}
