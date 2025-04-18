package model

import "time"

type Topic struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Destinations []string  `json:"destinations"`
	CreatedAt    time.Time `json:"created_at"`
}
