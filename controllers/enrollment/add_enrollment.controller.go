package enrollment

import (
	"errors"
	"final-project-prakerja-golang-batch-11/configs"
	"final-project-prakerja-golang-batch-11/models/database"
	"final-project-prakerja-golang-batch-11/models/response"
	"final-project-prakerja-golang-batch-11/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddEnrollmentController(context echo.Context) error {
	enrollment := new(database.Enrollment)
	context.Bind(&enrollment)

	// Validation
	if err := context.Validate(enrollment); err != nil {
		return context.JSON(http.StatusBadRequest, response.Base{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// Check Course ID
	course := new(database.Course)
	resultCourseCheck := configs.DB.Where("id = ?", enrollment.CourseID).First(&course)

	if resultCourseCheck.Error != nil {
		return context.JSON(http.StatusBadRequest, response.Base{
			Status:  false,
			Message: "Course ID Not Valid",
			Data:    nil,
		})
	}

	enrollment.UserID = utils.GetUserId(context)

	resultDataExists := configs.DB.
		Where("user_id = ? AND course_id = ?", enrollment.UserID, enrollment.CourseID).
		First(&enrollment)

	if errors.Is(resultDataExists.Error, gorm.ErrRecordNotFound) {
		resultAddEnrollment := configs.DB.Create(&enrollment)

		if resultAddEnrollment.Error == nil {
			return context.JSON(http.StatusOK, response.Base{
				Status:  true,
				Message: "Success Add Enrollment",
				Data:    nil,
			})
		}
	}

	if resultDataExists.Error == nil {
		return context.JSON(http.StatusConflict, response.Base{
			Status:  false,
			Message: "Enrollment already exists",
			Data:    nil,
		})
	}

	return context.JSON(http.StatusInternalServerError, response.Base{
		Status:  false,
		Message: "Server Error",
		Data:    nil,
	})
}
