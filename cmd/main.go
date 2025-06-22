package main

// main.go
// @title           Task Manager API
// @version         1.0.1
// @description     API documentation for the workmate task manager.
// @termsOfService  http://swagger.io/terms/
//
// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   gregorydyuldin@gmail.com
//
// @license.name    MIT
// @license.url     https://opensource.org/licenses/MIT
//
// @host            localhost:8080
// @BasePath        /api/v1

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	_ "workmate-test/cmd/docs"
	httpHandler "workmate-test/internal/handler/http"
	"workmate-test/internal/service"
	"workmate-test/internal/task"
)

func main() {
	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	manager := task.NewTaskManager(logger)
	svc := service.NewTaskService(manager)
	httpSrvHandler := httpHandler.NewHTTPHandler(svc, logger)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", "localhost", "8080"),
		Handler: httpSrvHandler,
	}

	sigChan := make(chan os.Signal, 1)
	done := make(chan error, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		done <- fmt.Errorf("received signal: %v", sig)
	}()

	go func() {
		_ = logger.Log("msg", "starting HTTP server", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			done <- fmt.Errorf("http server error: %w", err)
		}
	}()

	err := <-done
	_ = logger.Log("msg", "shutting down", "reason", err)

	// Optional: Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
