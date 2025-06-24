package model

import "time"

type Task struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAT time.Time `json:"updated_at"`
	Result    string    `json:"result,omitempty"`
}
