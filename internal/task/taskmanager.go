package task

import (
	"errors"
	"fmt"
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"math/rand"
	"sync"
	"time"
)

const minSecs = 180
const maxSecs = 300

type TaskManager struct {
	tasks map[string]*Task
	log   log.Logger
	mu    sync.RWMutex
}

func NewTaskManager(logger log.Logger) *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*Task),
		log:   logger,
	}
}

func (m *TaskManager) CreateTask() string {
	id := uuid.New().String()
	task := &Task{
		ID:        id,
		CreatedAt: time.Now(),
		Status:    statusPending,
	}
	m.mu.Lock()
	m.tasks[id] = task
	m.mu.Unlock()

	go m.processTask(task)
	return task.ID
}

func (m *TaskManager) GetTask(id string) (*Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	task, ok := m.tasks[id]
	if !ok {
		err := errors.New("wrong or non-existent id")
		_ = m.log.Log("error", fmt.Sprintf("failed to get task: %v", err))
		return nil, fmt.Errorf("failed to get task: %w", err)
	}
	return task, nil
}

func (m *TaskManager) DeleteTask(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.tasks[id]
	if !ok {
		err := errors.New("wrong or non-existent id")
		_ = m.log.Log("error", fmt.Sprintf("failed to delete task: %v", err))
		return fmt.Errorf("failed to delete task: %w", err)
	}
	delete(m.tasks, id)
	return nil
}

// processTask simulates a high load task that runs for a specifies amount of time
func (m *TaskManager) processTask(task *Task) {
	m.mu.Lock()
	task.Status = statusInProgress
	m.mu.Unlock()

	totalSecs := minSecs + rand.Intn(maxSecs-minSecs+1)
	delay := time.Duration(totalSecs) * time.Second
	time.Sleep(delay)

	m.mu.Lock()
	task.Status = statusDone
	now := time.Now()
	task.CompletedAt = &now
	task.Result = "Task completed successfully"
	m.mu.Unlock()
}
