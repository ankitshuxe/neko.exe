.PHONY: build run test clean

APP_NAME=neko

build:
	@echo "Building $(APP_NAME)..."
	go build -v -o bin/$(APP_NAME) .

run: build
	@echo "Running $(APP_NAME)..."
	./bin/$(APP_NAME)

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning up..."
	rm -rf bin/
