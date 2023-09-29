package database

type Course struct {
	ID          uint   `gorm:"type:int(11);primaryKey;autoIncrement" json:"id"`
	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Organizer   string `gorm:"type:varchar(255);not null" json:"organizer"`
	Price       int    `gorm:"type:int(11);default:0;not null" json:"price"`
}
