package models

import "time"

type LoginRecord struct {
	ID         int64
	User       User
	LoginTime  time.Time
	LogoutTime time.Time
}
