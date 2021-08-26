package models

import "time"

type Task struct {
	TaskID    int64     `json:"taskid"`
	UserID    string    `json:"userid"`
	Title     string    `json:"title"`
	Due       time.Time `json:"due"`
	Completed bool      `json:"completed"`
	Notes     string    `json:"notes"`
	Tags      []string  `json:"tags"`
}
