package dto

type CreateReservationRequest struct {
	ZoneID       uint   `json:"zone_id" validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required,min=3,max=15"`
}

type UpdateReservationStatusRequest struct {
	Status string `json:"status" validate:"required"`
}
