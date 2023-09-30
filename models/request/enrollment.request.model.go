package request

type EnrollmentRequest struct {
	CourseID uint `json:"courseId" validate:"required"`
}
