package decoder

import (
	"context"
	"errors"
	"net/http"
	"workmate-test/internal/handler/http/schemas"
)

// DecodeNewTaskRequest is a dummy function to comply to go-kits needs
func DecodeNewTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return schemas.MakeTaskRequest{}, nil
}

func DecodeGetTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	query := r.URL.Query()
	id := query.Get("id")

	if id == "" {
		return nil, errors.New("missing id")
	}

	req := schemas.GetTaskRequest{ID: id}
	return req, nil
}

func DecodeDeleteTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	query := r.URL.Query()
	id := query.Get("id")

	if id == "" {
		return nil, errors.New("missing id")
	}

	req := schemas.DeleteTaskRequest{ID: id}
	return req, nil
}
