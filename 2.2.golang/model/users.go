package model




type Users struct {
	Id       int    `gorm:"type:int;primary_key"`
	Username string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"type:varchar(255);not null"`
}