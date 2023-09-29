package database

type User struct {
	Id          uint   `gorm:"type:int(11);primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);not null"`
	Email       string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password    string `gorm:"type:varchar(255);not null"`
	PhoneNumber string `gorm:"type:varchar(15);not null"`
}
