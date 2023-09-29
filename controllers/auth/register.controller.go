package auth

import (
	"final-project-prakerja-golang-batch-11/configs"
	"final-project-prakerja-golang-batch-11/middleware"
	"final-project-prakerja-golang-batch-11/models/database"
	"final-project-prakerja-golang-batch-11/models/request"
	"final-project-prakerja-golang-batch-11/models/response"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(context echo.Context) error {
	request := new(request.RegisterRequest)
	context.Bind(&request)

	if err := context.Validate(request); err != nil {
		return context.JSON(http.StatusBadRequest, response.Base{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	request.Password = string(hashPassword)

	user := new(database.User)
	user.MapFromRegister(*request)

	result := configs.DB.Create(&user)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), fmt.Sprintf("Duplicate entry '%s' for key 'idx_users_email", request.Email)) {
			return context.JSON(http.StatusConflict, response.Base{
				Status:  false,
				Message: "User with that email already exists",
				Data:    nil,
			})
		} else {
			return context.JSON(http.StatusInternalServerError, response.Base{
				Status:  false,
				Message: "Failed Register",
				Data:    nil,
			})
		}
	}

	authResponse := response.Auth{
		Id:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Token:       middleware.GenerateTokenJWT(user.ID),
	}

	return context.JSON(http.StatusOK, response.Base{
		Status:  true,
		Message: "Success Register",
		Data:    authResponse,
	})
}
