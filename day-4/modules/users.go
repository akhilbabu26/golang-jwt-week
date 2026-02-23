package modules

import "gorm.io/gorm"

type Users struct{
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	Role string `json:"role"`
}


