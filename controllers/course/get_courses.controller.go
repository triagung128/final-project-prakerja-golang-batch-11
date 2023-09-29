package course

import (
	"final-project-prakerja-golang-batch-11/configs"
	"final-project-prakerja-golang-batch-11/models/database"
	"final-project-prakerja-golang-batch-11/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCoursesController(context echo.Context) error {
	courses := new([]database.Course)

	result := configs.DB.Find(&courses)

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
		Data:    courses,
	})
}
