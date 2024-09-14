package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Title               string  `json:"title"`
	Description         string  `json:"description"`
	JobType             string  `json:"job_type"`     // e.g., Full-time, Part-time, Internship
	Location            string  `json:"location"`     // e.g., On-site, Remote, Hybrid
	Eligibility         string  `json:"eligibility"`  // e.g., Degree requirements, cutoff marks
	Compensation        string  `json:"compensation"` // e.g., Salary or stipend details
	ApplicationDeadline string  `json:"application_deadline"`
	CompanyID           uint    `json:"company_id"` // Foreign key to Company
	Company             Company `gorm:"foreignKey:CompanyID"`
}
