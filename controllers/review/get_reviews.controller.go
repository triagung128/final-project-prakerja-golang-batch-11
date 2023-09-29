package review

import (
	"final-project-prakerja-golang-batch-11/configs"
	"final-project-prakerja-golang-batch-11/models/database"
	"final-project-prakerja-golang-batch-11/models/response"
	"final-project-prakerja-golang-batch-11/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetReviewsController(context echo.Context) error {
	reviews := new([]database.Review)
	userId := utils.GetUserId(context)

	result := configs.DB.
		InnerJoins("User").
		InnerJoins("Course").
		Where("user_id = ?", userId).
		Find(&reviews)

	if result.Error != nil {
		return context.JSON(http.StatusInternalServerError, response.Base{
			Status:  false,
			Message: "Server Error",
			Data:    nil,
		})
	}

	return context.JSON(http.StatusOK, response.Base{
		Status:  true,
		Message: "Success",
		Data:    reviews,
	})
}
