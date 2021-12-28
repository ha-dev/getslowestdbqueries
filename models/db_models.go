package models

import "gorm.io/gorm"

// Book struct to describe book object.
type Querys struct {
	gorm.Model
	Query       string `db:"query" json:"Query"`
	Time_to_run string `db:"time_to_run" json:"Time_to_run"`
}

// BookAttrs struct to describe book attributes.
type ResulyQuery struct {
	ID          uint64 `json:"id"`
	Query       string `json:"query"`
	Time_to_run int    `json:"time_to_run"`
}
