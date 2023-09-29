package database

import "time"

type Enrollment struct {
	ID        uint      `gorm:"type:int(11);not null" json:"id"`
	UserID    uint      `gorm:"type:int(11);not null" json:"userId"`
	User      User      `json:"user"`
	CourseID  uint      `gorm:"type:int(11);not null" json:"courseId" validate:"required"`
	Course    Course    `json:"course"`
	CreatedAt time.Time `json:"createdAt"`
}
