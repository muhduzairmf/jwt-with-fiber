package models

import "time"

type Profile struct {
	ID uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Birthday string `json:"birthday"`
	Website string `json:"website"`
	UserId uint `json:"userId"`
	User User `gorm:"foreignKey"`
}
