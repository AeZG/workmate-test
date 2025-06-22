package encoder

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"net/http"
	"workmate-test/internal/handler/http/schemas"
)

// EncodeResponse encodes the responses into an HTTP response
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	switch res := response.(type) {
	case schemas.MakeTaskResponse:
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(res)
	case schemas.GetTaskResponse:
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(res)
	case schemas.DeleteTaskResponse:
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(res)
	default:
		return fmt.Errorf("unknown response type: %T", response)
	}
}

// EncodeError encodes errors into an HTTP error response
func EncodeError(logger log.Logger) func(_ context.Context, err error, w http.ResponseWriter) {
	return func(_ context.Context, err error, w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		var code int
		var msg string

		switch {
		case isValidationError(err):
			code = http.StatusBadRequest
			msg = err.Error()
		default:
			code = http.StatusInternalServerError
			// Log the actual error but return a generic message
			_ = logger.Log("error", fmt.Sprintf("Internal server error: %v", err))
			msg = http.StatusText(http.StatusInternalServerError)
		}

		w.WriteHeader(code)
		if err := json.NewEncoder(w).Encode(schemas.ErrorResponse{Error: msg}); err != nil {
			_ = logger.Log("error", fmt.Sprintf("Failed to encode error response: %v", err))
		}
	}
}

// isValidationError checks if the error is due to validation difficulties
func isValidationError(err error) bool {
	return true
}
