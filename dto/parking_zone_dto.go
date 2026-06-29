package dto

type CreateParkingZoneRequest struct {
	Name          string  `json:"name" validate:"required,min=3,max=100"`
	Type          string  `json:"type" validate:"required"`
	TotalCapacity int     `json:"total_capacity" validate:"required,gt=0"`
	PricePerHour  float64 `json:"price_per_hour" validate:"required,gt=0"`
}

type UpdateParkingZoneRequest struct {
	Name          string  `json:"name" validate:"required,min=3,max=100"`
	Type          string  `json:"type" validate:"required"`
	TotalCapacity int     `json:"total_capacity" validate:"required,gt=0"`
	PricePerHour  float64 `json:"price_per_hour" validate:"required,gt=0"`
}
