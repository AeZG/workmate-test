package http

import (
	"encoding/json"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"workmate-test/internal/handler/http/decoder"
	"workmate-test/internal/handler/http/encoder"
	"workmate-test/internal/service"
)

const (
	apiPrefix = "/api"
	v1Prefix  = apiPrefix + "/v1"
)

// Create a new HTTP handler
func NewHTTPHandler(svc service.TaskService, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc(apiPrefix, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]interface{}{
			"versions":      []string{"v1"},
			"documentation": v1Prefix + "/docs/index.html",
		})
		if err != nil {
			_ = level.Error(logger).Log("msg", "failed to encode API info", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	registerV1Routes(logger, r, svc)

	return r
}

// registerV1routes registers the routes for v1 of the API
func registerV1Routes(logger log.Logger, router *mux.Router, svc service.TaskService) {
	v1 := router.PathPrefix(v1Prefix).Subrouter()

	v1.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL(v1Prefix+"/docs/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("list"),
		httpSwagger.DomID("swagger-ui"),
	))

	v1.Methods("POST").Path("/task").Handler(kitHttp.NewServer(
		MakeNewTaskEndpoint(svc),
		decoder.DecodeNewTaskRequest,
		encoder.EncodeResponse,
		kitHttp.ServerErrorEncoder(encoder.EncodeError(logger)),
	))

	v1.Methods("GET").Path("/task").Handler(kitHttp.NewServer(
		MakeGetTaskEndpoint(svc),
		decoder.DecodeGetTaskRequest,
		encoder.EncodeResponse,
		kitHttp.ServerErrorEncoder(encoder.EncodeError(logger)),
	))

	v1.Methods("DELETE").Path("/task").Handler(kitHttp.NewServer(
		MakeDeleteTaskEndpoint(svc),
		decoder.DecodeDeleteTaskRequest,
		encoder.EncodeResponse,
		kitHttp.ServerErrorEncoder(encoder.EncodeError(logger)),
	))
}
