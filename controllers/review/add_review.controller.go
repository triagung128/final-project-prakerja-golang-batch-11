package review

import (
	"errors"
	"final-project-prakerja-golang-batch-11/configs"
	"final-project-prakerja-golang-batch-11/models/database"
	"final-project-prakerja-golang-batch-11/models/request"
	"final-project-prakerja-golang-batch-11/models/response"
	"final-project-prakerja-golang-batch-11/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddReviewContoller(context echo.Context) error {
	request := new(request.ReviewRequest)
	context.Bind(&request)

	if err := context.Validate(request); err != nil {
		return context.JSON(http.StatusBadRequest, response.Base{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	review := new(database.Review)
	review.MapFromRequest(*request)

	review.UserID = utils.GetUserId(context)

	enrollment := new(database.Enrollment)
	resultEnrollmentCheck := configs.DB.
		Where("user_id = ? AND course_id = ?", review.UserID, review.CourseID).
		First(&enrollment)

	if resultEnrollmentCheck.Error != nil {
		return context.JSON(http.StatusMethodNotAllowed, response.Base{
			Status:  false,
			Message: "Not allowed for review",
			Data:    nil,
		})
	}

	resultDataExists := configs.DB.
		Where("user_id = ? AND course_id = ?", review.UserID, review.CourseID).
		First(&review)

	if errors.Is(resultDataExists.Error, gorm.ErrRecordNotFound) {
		resultAddReview := configs.DB.Create(&review)

		if resultAddReview.Error == nil {
			return context.JSON(http.StatusOK, response.Base{
				Status:  true,
				Message: "Success Add Review Course",
				Data:    nil,
			})
		}
	}

	if resultDataExists.Error == nil {
		return context.JSON(http.StatusConflict, response.Base{
			Status:  false,
			Message: "Review with that course already exists",
			Data:    nil,
		})
	}

	return context.JSON(http.StatusInternalServerError, response.Base{
		Status:  false,
		Message: "Server Error",
		Data:    nil,
	})
}
