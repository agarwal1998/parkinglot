package responseDto

type VehicleInfoByColorOp struct {
	RegistrationNo string
	Color          string
}

type VehicleInfoByRegistrationNoOp struct {
	RegistrationNo string
	Color          string
	SLotId         int
}

type SlotStatus struct {
	No                    int
	Status                string
	VehicleRegistrationNo string
}
