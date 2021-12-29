package models

import "gorm.io/gorm"

type Pagination struct {
	Totalrecord uint64
}

// Querys struct to describe Querys object.
type Querys struct {
	gorm.Model
	Query       string `db:"query" json:"Query"`
	Time_to_run string `db:"time_to_run" json:"Time_to_run"`
}

// ResulyQuery struct to describe Querys attributes.
type ResulyQuery struct {
	ID          uint64 `json:"id"`
	Query       string `json:"query"`
	Time_to_run int    `json:"time_to_run"`
}

type ResulyQueryWithPagination struct {
	Pagination
	ResulyQuery
}
