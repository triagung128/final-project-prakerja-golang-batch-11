package database

import "final-project-prakerja-golang-batch-11/models/request"

type User struct {
	ID          uint   `gorm:"type:int(11);primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Email       string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password    string `gorm:"type:varchar(255);not null" json:"password"`
	PhoneNumber string `gorm:"type:varchar(15);not null" json:"phoneNumber"`
}

func (user *User) MapFromRegister(request request.RegisterRequest) {
	user.Name = request.Name
	user.Email = request.Email
	user.Password = request.Password
	user.PhoneNumber = request.PhoneNumber
}

func (user *User) MapFromLogin(request request.LoginRequest) {
	user.Email = request.Email
	user.Password = request.Password
}
