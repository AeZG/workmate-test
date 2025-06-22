package schemas

import "time"

// These are the schemas required for creating responses and requests

type MakeTaskRequest struct{}

type MakeTaskResponse struct {
	Message string `json:"message"`
}

type GetTaskRequest struct {
	ID string `json:"ID" validate:"required"`
}

type GetTaskResponse struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
	Status      string     `json:"status"`
	Result      string     `json:"result"`
	ElapsedTime string     `json:"elapsed_time"`
}

type DeleteTaskRequest struct {
	ID string `json:"id" validate:"required"`
}

type DeleteTaskResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
