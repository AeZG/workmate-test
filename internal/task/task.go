package task

import "time"

type TaskStatus string

const (
	statusPending    TaskStatus = "pending"
	statusInProgress TaskStatus = "in progress"
	statusDone       TaskStatus = "done"
)

type Task struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
	Status      TaskStatus `json:"status"`
	Result      string     `json:"result"`
}
