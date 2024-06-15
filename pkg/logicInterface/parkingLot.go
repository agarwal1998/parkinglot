package logicInterface

import (
	"yo/pkg/entity"
	"yo/pkg/entity/responseDto"
)

type ParkingLotInterface interface {
	Create(count int)
	Insert(registrationNo, color string) (bool, entity.Slot)
	Remove(slotId int) entity.Vehicle
	GetByRegistrationNo(registrationNo string) responseDto.VehicleInfoByRegistrationNoOp
	GetByColor(color string) []responseDto.VehicleInfoByColorOp
	Status() []responseDto.SlotStatus
}
