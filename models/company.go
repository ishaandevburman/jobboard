package models

import "gorm.io/gorm"

type Company struct {
    gorm.Model
    Name        string `json:"name"`
    ContactEmail string `json:"contact_email"`
    ContactPerson string `json:"contact_person"`
    PhoneNumber  string `json:"phone_number"`
    Description  string `json:"description"`
    LogoURL      string `json:"logo_url"`
}
