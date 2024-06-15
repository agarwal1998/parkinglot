package logic

import (
	"yo/pkg/entity"
	"yo/pkg/entity/responseDto"
	repository "yo/pkg/repository/interface"
)

type ParkingLot struct {
	parkingRepo repository.ParkingRepoInterface
}

func NewParkingLot(parkingRepo repository.ParkingRepoInterface) *ParkingLot {
	return &ParkingLot{parkingRepo}
}

func (pl *ParkingLot) Create(count int) {
	pl.parkingRepo.Create(count)
}

func (pl *ParkingLot) Insert(registrationNo, color string) (bool, entity.Slot) {
	vehicle := entity.Vehicle{registrationNo, color, nil}
	return pl.parkingRepo.Insert(vehicle)
}

func (pl *ParkingLot) Remove(slotId int) entity.Vehicle {
	return pl.parkingRepo.Remove(slotId)
}

func (pl *ParkingLot) GetByRegistrationNo(registrationNo string) responseDto.VehicleInfoByRegistrationNoOp {
	data := pl.parkingRepo.GetVehicleByRegistrationNo(registrationNo)
	return responseDto.VehicleInfoByRegistrationNoOp{data.RegistrationNo, data.Color, data.Slot.No}
}

func (pl *ParkingLot) GetByColor(color string) []responseDto.VehicleInfoByColorOp {
	vehicles := pl.parkingRepo.GetByColor(color)
	ans := make([]responseDto.VehicleInfoByColorOp, 0, len(vehicles))
	for _, v := range vehicles {
		ans = append(ans, responseDto.VehicleInfoByColorOp{v.RegistrationNo, v.Color})
	}
	return ans
}

func (pl *ParkingLot) Status() []responseDto.SlotStatus {
	data := pl.parkingRepo.GetAllSlots()
	return formatData(data)
}

func formatData(slots []entity.Slot) []responseDto.SlotStatus {
	result := make([]responseDto.SlotStatus, 0, len(slots))
	for _, slot := range slots {
		slotStatus := responseDto.SlotStatus{
			No:     slot.No,
			Status: slot.Status,
		}
		if slot.Vehicle != nil {
			slotStatus.VehicleRegistrationNo = slot.Vehicle.RegistrationNo
		}
		result = append(result, slotStatus)
	}
	return result
}
