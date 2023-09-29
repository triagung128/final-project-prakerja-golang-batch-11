package auth

import (
	"errors"
	"final-project-prakerja-golang-batch-11/configs"
	"final-project-prakerja-golang-batch-11/middleware"
	"final-project-prakerja-golang-batch-11/models/database"
	"final-project-prakerja-golang-batch-11/models/request"
	"final-project-prakerja-golang-batch-11/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginController(context echo.Context) error {
	request := new(request.LoginRequest)
	context.Bind(&request)

	if err := context.Validate(request); err != nil {
		return context.JSON(http.StatusBadRequest, response.Base{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	user := new(database.User)
	user.MapFromLogin(*request)

	result := configs.DB.Where("email = ?", user.Email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return context.JSON(http.StatusUnauthorized, response.Base{
			Status:  false,
			Message: "Wrong Email or Password",
			Data:    nil,
		})
	}

	if result.Error != nil {
		return context.JSON(http.StatusInternalServerError, response.Base{
			Status:  false,
			Message: "Server Error",
			Data:    nil,
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return context.JSON(http.StatusUnauthorized, response.Base{
			Status:  false,
			Message: "Wrong Email or Password",
			Data:    nil,
		})
	}

	authResponse := response.Auth{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Token:       middleware.GenerateTokenJWT(user.Id),
	}

	return context.JSON(http.StatusOK, response.Base{
		Status:  true,
		Message: "Success Login",
		Data:    authResponse,
	})

}
