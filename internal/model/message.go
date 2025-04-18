package model

import "time"

type Message struct {
	ID          string    `json:"id"`
	TopicID     string    `json:"topic_id"`
	Body        string    `json:"body"`
	URL         string    `json:"url"`
	Delay       int64     `json:"delay"`
	Retries     int       `json:"retries"`
	ScheduledAt time.Time `json:"scheduled_at"`
	CreatedAt   time.Time `json:"created_at"`
}
