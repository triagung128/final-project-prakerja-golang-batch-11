package database

import "time"

type Review struct {
	ID        uint      `gorm:"type:int(11);not null" json:"id"`
	UserID    uint      `gorm:"type:int(11);not null" json:"userId"`
	User      User      `json:"user"`
	CourseID  uint      `gorm:"type:int(11);not null" json:"courseId" validate:"required"`
	Course    Course    `json:"course"`
	Review    string    `gorm:"type:text;not null" json:"review" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
}
