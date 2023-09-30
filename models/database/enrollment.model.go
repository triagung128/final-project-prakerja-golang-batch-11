package database

import (
	"final-project-prakerja-golang-batch-11/models/request"
	"time"
)

type Enrollment struct {
	ID        uint      `gorm:"type:int(11);not null" json:"id"`
	UserID    uint      `gorm:"type:int(11);not null" json:"-"`
	User      User      `json:"user"`
	CourseID  uint      `gorm:"type:int(11);not null" json:"-"`
	Course    Course    `json:"course"`
	CreatedAt time.Time `json:"createdAt"`
}

func (enrollment *Enrollment) MapFromRequest(request request.EnrollmentRequest) {
	enrollment.CourseID = request.CourseID
}
