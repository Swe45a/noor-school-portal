package domain

import "time"

type Teacher struct {
	ID                  string
	IDNumber            string
	EmployeeNumber      string
	FullName            string
	AccountUsername     string
	AccountPasswordHash *string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
