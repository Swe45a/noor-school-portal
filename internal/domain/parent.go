package domain

import "time"

type Parent struct {
	ID                  string
	ParentID            string
	FullName            string
	AccountUsername     string
	AccountPasswordHash *string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
