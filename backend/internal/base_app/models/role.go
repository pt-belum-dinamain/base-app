package models

import "time"

type Role struct {
	ID       int64
	RoleName string

	CreatedAt time.Time
	CreatedBy User
	UpdatedAt time.Time
	UpdatedBy User
}
