package models

import "time"

type User struct {
	ID 		   uint `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"createdAt"`
	Email 	   string `json:"email" gorm:"unique"`
	Password   string `json:"password"`
}
