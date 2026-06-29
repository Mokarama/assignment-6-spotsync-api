package models

import "gorm.io/gorm"

type ParkingZone struct {
	gorm.Model

	Name          string  `gorm:"type:varchar(100);not null"`
	Type          string  `gorm:"type:varchar(30);not null"`
	TotalCapacity int     `gorm:"not null"`
	PricePerHour  float64 `gorm:"type:decimal(10,2);not null"`

	AvailableSpots int `gorm:"-" json:"available_spots"`
}