package database

import "time"

type Review struct {
	Id        uint   `gorm:"type:int(11);not null"`
	UserID    uint   `gorm:"type:int(11);not null"`
	CourseId  uint   `gorm:"type:int(11);not null"`
	Review    string `gorm:"type:text;not null"`
	CreatedAt time.Time
}
