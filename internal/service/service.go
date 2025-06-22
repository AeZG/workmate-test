package service

import (
	"fmt"
	"time"
	"workmate-test/internal/task"
)

// --- This layer is used for business logic for better separation of code

type TaskService interface {
	CreateTask() string
	GetTask(id string) (*task.Task, string, error)
	DeleteTask(id string) error
}

type taskService struct {
	tm *task.TaskManager
}

func NewTaskService(taskManager *task.TaskManager) TaskService {
	return &taskService{tm: taskManager}
}

func (s *taskService) CreateTask() string {
	taskID := s.tm.CreateTask()
	return taskID
}

func (s *taskService) GetTask(id string) (*task.Task, string, error) {
	taskInfo, err := s.tm.GetTask(id)
	if err != nil {
		return nil, "", err
	}
	elapsedTime := getElapsedTime(taskInfo.CreatedAt, taskInfo.CompletedAt)
	return taskInfo, elapsedTime, nil
}

func (s *taskService) DeleteTask(id string) error {
	err := s.tm.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func getElapsedTime(createdAt time.Time, completedAt *time.Time) string {
	if completedAt == nil {
		elapsedTime := fmt.Sprintf("%.2fs", time.Since(createdAt).Seconds())
		return elapsedTime
	}
	return fmt.Sprintf("%.2fs", completedAt.Sub(createdAt).Seconds())
}
