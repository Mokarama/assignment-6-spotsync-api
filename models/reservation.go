package models

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model

	UserID uint `gorm:"not null"`
	User   User `gorm:"foreignKey:UserID"`

	ZoneID uint        `gorm:"not null"`
	Zone   ParkingZone `gorm:"foreignKey:ZoneID"`

	LicensePlate string `gorm:"type:varchar(15);not null"`
	Status       string `gorm:"type:varchar(20);default:active"`
}
