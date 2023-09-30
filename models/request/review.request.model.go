package request

type ReviewRequest struct {
	CourseID uint   `gorm:"type:int(11);not null" json:"courseId" validate:"required"`
	Review   string `gorm:"type:text;not null" json:"review" validate:"required"`
}
