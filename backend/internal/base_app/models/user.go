package models

import "time"

type User struct {
	ID        int64
	Name      string
	Password  string
	LastLogin time.Time
	Roles     []Role

	CreatedAt time.Time
	CreatedBy User
	UpdatedAt time.Time
	UpdatedBy User
}
