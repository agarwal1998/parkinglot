package entity

import (
	"container/heap"
	"yo/constants"
)

type FloorInterface interface {
	GetEmptySlot() int
	GetSlots() []Slot
	GetVehicleAt(slotId int) *Vehicle
	ResetSlot(slotId int)
	AddVehicle(slotId int, vehicle *Vehicle) *Slot
}

type Floor struct {
	Slots         []Slot
	EmptySlotHeap IntHeap
}

func (f *Floor) GetEmptySlot() int {
	if len(f.EmptySlotHeap) == 0 {
		return -1
	}
	return heap.Pop(&f.EmptySlotHeap).(int)
}

func (f *Floor) GetSlots() []Slot {
	return f.Slots
}

func (f *Floor) GetVehicleAt(slotId int) *Vehicle {
	return f.Slots[slotId-1].Vehicle
}

func (f *Floor) ResetSlot(slotId int) {
	f.Slots[slotId-1] = Slot{slotId, nil, constants.SlotStatus.EMPTY}
	heap.Push(&f.EmptySlotHeap, slotId)
}
func (f *Floor) AddVehicle(slotId int, vehicle *Vehicle) *Slot {
	f.Slots[slotId-1] = Slot{slotId, vehicle, constants.SlotStatus.FILLED}
	return &f.Slots[slotId-1]
}
