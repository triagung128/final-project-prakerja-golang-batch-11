package routes

import (
	authController "final-project-prakerja-golang-batch-11/controllers/auth"
	courseController "final-project-prakerja-golang-batch-11/controllers/course"
	enrollmentController "final-project-prakerja-golang-batch-11/controllers/enrollment"
	jwtMiddleware "final-project-prakerja-golang-batch-11/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	e.POST("/register", authController.RegisterController)
	e.POST("/login", authController.LoginController)
	e.GET("/courses", courseController.GetCoursesController)

	eJWT := e.Group("")
	eJWT.Use(echojwt.WithConfig(jwtMiddleware.GetConfig()))
	eJWT.GET("/enrollments", enrollmentController.GetEnrollmentsController)
	eJWT.POST("/enrollments", enrollmentController.AddEnrollmentController)
}
