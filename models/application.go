package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	StudentID   uint    `json:"student_id"`
	JobID       uint    `json:"job_id"`
	Status      string  `json:"status"`
	CoverLetter string  `json:"cover_letter"`
	ResumeURL   string  `json:"resume_url"`
	Student     Student `gorm:"foreignKey:StudentID"`
	Job         Job     `gorm:"foreignKey:JobID"`
}
