package configs

import "final-project-prakerja-golang-batch-11/models/database"

func InitMigrate() {
	DB.AutoMigrate(&database.User{})
	DB.AutoMigrate(&database.Course{})
	DB.AutoMigrate(&database.Enrollment{})
	DB.AutoMigrate(&database.Review{})
}
