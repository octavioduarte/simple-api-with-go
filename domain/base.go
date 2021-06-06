package domain

import "time"

type Base struct {
	ID        string
	Active    bool
	CreatedAt time.Timer
	UpdatedAt time.Timer
	DeletedAt time.Timer
}
