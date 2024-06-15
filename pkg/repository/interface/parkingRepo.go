package repository

import (
	"yo/pkg/entity"
)

type ParkingRepoInterface interface {
	Insert(vehicle entity.Vehicle) (bool, entity.Slot)
	Create(count int)
	Remove(slotId int) entity.Vehicle
	GetAllSlots() []entity.Slot
	GetVehicleByRegistrationNo(registrationNo string) *entity.Vehicle
	GetByColor(color string) []*entity.Vehicle
}
