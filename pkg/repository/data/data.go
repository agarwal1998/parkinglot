package repository

import (
	"container/heap"
	"yo/constants"
	"yo/pkg/entity"
)

type ParkingRepo struct {
	Floor                    entity.FloorInterface
	ColorVehicleMap          map[string][]*entity.Vehicle
	RegistrationNoVehicleMap map[string]*entity.Vehicle
}

func NewParkingRepo() *ParkingRepo {
	return &ParkingRepo{ColorVehicleMap: make(map[string][]*entity.Vehicle, 0), RegistrationNoVehicleMap: make(map[string]*entity.Vehicle, 0)}
}

func (pf *ParkingRepo) Create(count int) {
	slots := make([]entity.Slot, 0, count)
	emptySlotHeap := make(entity.IntHeap, 0, count)
	for i := 0; i < count; i++ {
		slots = append(slots, entity.Slot{Status: constants.SlotStatus.EMPTY, No: i + 1})
		emptySlotHeap = append(emptySlotHeap, i+1)
	}
	heap.Init(&emptySlotHeap)
	pf.Floor = &entity.Floor{slots, emptySlotHeap}
}

func (pf *ParkingRepo) Insert(vehicle entity.Vehicle) (bool, entity.Slot) {
	floor := pf.Floor
	if floor == nil {
		return false, entity.Slot{}
	}
	slotId := floor.GetEmptySlot()
	if slotId == -1 {
		return false, entity.Slot{}
	}
	newSlot := floor.AddVehicle(slotId, &vehicle)
	vehicle.Slot = newSlot
	pf.AddVehicle(vehicle)
	return true, *newSlot
}

func (pf *ParkingRepo) Remove(slotId int) entity.Vehicle {
	vehicle := pf.Floor.GetVehicleAt(slotId)
	pf.RemoveVehicle(*vehicle)
	pf.Floor.ResetSlot(slotId)
	return *vehicle
}

func (pf *ParkingRepo) GetAllSlots() []entity.Slot {
	return pf.Floor.GetSlots()
}

func (pf *ParkingRepo) GetVehicleByRegistrationNo(registrationNo string) *entity.Vehicle {
	return pf.RegistrationNoVehicleMap[registrationNo]
}

func (pf *ParkingRepo) GetByColor(color string) []*entity.Vehicle {
	return pf.ColorVehicleMap[color]
}

func (pf *ParkingRepo) AddVehicle(vehicle entity.Vehicle) {
	pf.ColorVehicleMap[vehicle.Color] = append(pf.ColorVehicleMap[vehicle.Color], &vehicle)
	pf.RegistrationNoVehicleMap[vehicle.RegistrationNo] = &vehicle
}

func (pf *ParkingRepo) RemoveVehicle(vehicle entity.Vehicle) {
	delete(pf.RegistrationNoVehicleMap, vehicle.RegistrationNo)
	vehiclesByColour := pf.ColorVehicleMap[vehicle.Color]
	newVehiclesByColour := make([]*entity.Vehicle, 0, len(vehiclesByColour)-1)
	for i := 0; i < len(vehiclesByColour); i++ {
		if vehiclesByColour[i].RegistrationNo != vehicle.RegistrationNo {
			newVehiclesByColour = append(newVehiclesByColour, vehiclesByColour[i])
		}
	}
	pf.ColorVehicleMap[vehicle.Color] = newVehiclesByColour
}
