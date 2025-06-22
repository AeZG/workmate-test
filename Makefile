# Project configuration
APP_NAME = workmate
MAIN_PKG = cmd/main.go
SWAG_OUTPUT = cmd/docs

.PHONY: all
all: run

.PHONY: run
run:
	go run $(MAIN_PKG)

.PHONY: build
build:
	go build -o $(APP_NAME) $(MAIN_PKG)

.PHONY: clean
clean:
	rm -f $(APP_NAME)

.PHONY: rebuild
rebuild: docs build