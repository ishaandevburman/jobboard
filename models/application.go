package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	StudentID   uint    `json:"student_id"` // Foreign key to Student
	JobID       uint    `json:"job_id"`     // Foreign key to Job
	Status      string  `json:"status"`     // e.g., Applied, Under Review, Shortlisted, Rejected
	CoverLetter string  `json:"cover_letter"`
	ResumeURL   string  `json:"resume_url"` // URL or path to student's resume
	Student     Student `gorm:"foreignKey:StudentID"`
	Job         Job     `gorm:"foreignKey:JobID"`
}
