package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"yo/constants"
	"yo/pkg/logicInterface"
)

type ApiHandler struct {
	ParkingLotLogic logicInterface.ParkingLotInterface
}

func NewApiHandler(parkingLotLogic logicInterface.ParkingLotInterface) *ApiHandler {
	return &ApiHandler{parkingLotLogic}
}

// create lot 5
// parking lot created
func (ah *ApiHandler) Create(w1, w2, w3 string) {
	count, err := strconv.Atoi(w2)
	if err != nil {
		fmt.Println("InvalidInput")
	}
	ah.ParkingLotLogic.Create(count)
	fmt.Println("parking lot created")
}

// park KA01AB1234 White
func (ah *ApiHandler) Insert(w1, w2, w3 string) {
	done, slot := ah.ParkingLotLogic.Insert(w1, w2)
	if done {
		fmt.Printf("Parked at slot %d\n", slot.No)
	} else {
		fmt.Println("Parking log is full")
	}
}

// leave Slot 2
// Leave KA03DE3434
func (ah *ApiHandler) Remove(w1, w2, w3 string) {
	count, err := strconv.Atoi(w2)
	if err != nil {
		fmt.Println("InvalidInput")
	}
	vehicle := ah.ParkingLotLogic.Remove(count)
	if vehicle.RegistrationNo != "" {
		//fmt.Println(fmt.Scanf("Leave %s", vehicle.RegistrationNo))
		fmt.Printf("Leave %v\n", vehicle.RegistrationNo)
	} else {
		fmt.Println("Slot Already Empty")
	}
}

// status
// Slot 1 -> KA01AB1234 Car parked
// Slot 2 -> KA05VB4567 Car parked
// Slot 3 -> KA04BC6734 Car parked
// Slot 4 -> Empty
// Slot 5 -> Empty
func (ah *ApiHandler) Status(w1, w2, w3 string) {
	data := ah.ParkingLotLogic.Status()
	for _, slotStatus := range data {
		if slotStatus.Status == constants.SlotStatus.FILLED {
			fmt.Printf("Slot %v -> %v Car parked\n", slotStatus.No, slotStatus.VehicleRegistrationNo)
		} else {
			fmt.Printf("Slot %v -> Empty\n", slotStatus.No)
		}
	}
}

// color White
// Car with White Color : KA01AB1234, KA04BC6734
func (ah *ApiHandler) GetByColor(w1, w2, w3 string) {
	data := ah.ParkingLotLogic.GetByColor(w1)
	registrationNos := make([]string, 0, len(data))
	for _, vehicle := range data {
		registrationNos = append(registrationNos, vehicle.RegistrationNo)
	}
	fmt.Printf("Car with %v Color : %v\n", w1, strings.Join(registrationNos, ","))
}

// registration KA04BC6734
// registration KA04BC6734 parked at slot 3
func (ah *ApiHandler) GetByRegistrationNo(w1, w2, w3 string) {
	data := ah.ParkingLotLogic.GetByRegistrationNo(w1)
	fmt.Printf("registration %v parked at slot %v\n", data.RegistrationNo, data.SLotId)
}
