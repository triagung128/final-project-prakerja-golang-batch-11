package database

import (
	"final-project-prakerja-golang-batch-11/models/request"
	"time"
)

type Review struct {
	ID        uint      `gorm:"type:int(11);not null" json:"id"`
	UserID    uint      `gorm:"type:int(11);not null" json:"-"`
	User      User      `json:"user"`
	CourseID  uint      `gorm:"type:int(11);not null" json:"-" validate:"required"`
	Course    Course    `json:"course"`
	Review    string    `gorm:"type:text;not null" json:"review" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

func (review *Review) MapFromRequest(request request.ReviewRequest) {
	review.CourseID = request.CourseID
	review.Review = request.Review
}
