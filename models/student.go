package models

import "gorm.io/gorm"

type Student struct {
    gorm.Model
    RegistrationNumber string `json:"registration_number" gorm:"unique"`
    Email              string `json:"email" gorm:"unique"`
    Password           string `json:"password"`
    Resume             string `json:"resume"`
}
