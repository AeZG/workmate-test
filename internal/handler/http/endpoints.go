package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"workmate-test/internal/handler/http/schemas"
	"workmate-test/internal/service"
)

const (
	invalidRequestTypeErrorMessage = "invalid request type"
)

// MakeNewTaskEndpoint handles creating a task
//
// @Summary      Create Task
// @Description  Creates a new task.
// @Tags         Task
// @Accept       json
// @Produce      json
// @Success      200          {object}  schemas.MakeTaskResponse
// @Failure      400          {object}  schemas.ErrorResponse
// @Failure      500          {object}  schemas.ErrorResponse
// @Router       /task [post]
func MakeNewTaskEndpoint(svc service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		taskID := svc.CreateTask()
		return schemas.MakeTaskResponse{
			Message: fmt.Sprintf("Task with ID %s has been created", taskID),
		}, nil
	}
}

// MakeGetTaskEndpoint handles getting task info
//
// @Summary      Get task info
// @Description  Returns all info for a specific task by ID.
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        id           query     string  true  "ID"
// @Success      200          {object}  schemas.GetTaskResponse
// @Failure      400          {object}  schemas.ErrorResponse
// @Failure      500          {object}  schemas.ErrorResponse
// @Router       /task [get]
func MakeGetTaskEndpoint(svc service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(schemas.GetTaskRequest)
		if !ok {
			return nil, errors.New(invalidRequestTypeErrorMessage)
		}

		taskInfo, elapsedTime, err := svc.GetTask(req.ID)

		if err != nil {
			return nil, err
		}

		return schemas.GetTaskResponse{
			ID:          taskInfo.ID,
			CreatedAt:   taskInfo.CreatedAt,
			CompletedAt: taskInfo.CompletedAt,
			Status:      string(taskInfo.Status),
			Result:      taskInfo.Result,
			ElapsedTime: elapsedTime,
		}, nil

	}
}

// MakeDeleteTaskEndpoint handles getting task info
//
// @Summary      Delete task
// @Description  Deletes task with specified ID.
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        id           query     string  true  "ID"
// @Success      200          {object}  schemas.DeleteTaskResponse
// @Failure      400          {object}  schemas.ErrorResponse
// @Failure      500          {object}  schemas.ErrorResponse
// @Router       /task [delete]
func MakeDeleteTaskEndpoint(svc service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(schemas.DeleteTaskRequest)

		if !ok {
			return nil, errors.New(invalidRequestTypeErrorMessage)
		}

		err = svc.DeleteTask(req.ID)
		if err != nil {
			return nil, err
		}

		return schemas.DeleteTaskResponse{
			Message: "Task Deleted successfully",
		}, nil
	}
}
