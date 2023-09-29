package database

type Course struct {
	Id          uint   `gorm:"type:int(11);primaryKey;autoIncrement"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Organizer   string `gorm:"type:varchar(255);not null"`
	Price       int    `gorm:"type:int(11);default:0;not null"`
}
