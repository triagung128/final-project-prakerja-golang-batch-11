package database

import "time"

type Enrollment struct {
	Id        uint `gorm:"type:int(11);not null"`
	UserID    uint `gorm:"type:int(11);not null"`
	Course    uint `gorm:"type:int(11);not null"`
	CreatedAt time.Time
}
