.PHONY: default run build test clean

APP_NAME = gopportunities
APP_VERSION = 0.1.0

# tasks
default: run

run:
	@go run main.go
build:
	@go build -o $(APP_NAME)  main.go
test:
	@go test ./ ...
clean:
	@rm -f $(APP_NAME)
