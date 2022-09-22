package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type Users struct {
	gorm.Model

	Employee_id int    `json:"employee_id" validate:"required" `
	Name        string `json:"name" validate:"required,alpha,startswith=A"`
	Lastname    string `json:"lastname" validate:"required"`
	Birthday    string `json:"birthday" validate:"required"`
	Age         int    `json:"age" validate:"required"`
	Email       string `json:"email,omitempty" validate:"required,email"`
	Tel         string `json:"tel" validate:"required" `

	// Pass      string `json:"pass" validate:"required,min=2,max=20"`

	//Bussiness string `json:"bussiness" validate:"required"`
	//Nameweb   string `json:"nameweb" validate:"required,min=2,max=30"`
	// Age uint8 `validate: "gte=0,lte=130 "`

}
