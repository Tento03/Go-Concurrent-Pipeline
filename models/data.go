package models

import "time"

type Data struct {
	ID        int
	Payload   string
	CreatedAt time.Time
}
