APP_BINARY=cquestApp

# Build binary
build:
	@echo "Building app..."
	CGO_ENABLED=0 go build -o _build/$(APP_BINARY) cmd/main.go
	@echo "Done!"

# Build and start the binary
run-build: build
	@echo "Starting app..."
	_build/$(APP_BINARY)
	@echo "Done!"

# Run project
run:
	@echo "Starting app..."
	go run cmd/main.go
	@echo "Done!"

# Dockerize
dockerize:
	@echo "Building app..."
	docker build -t alpeshjamgade/cpu-compare:$(TAG) .
	@echo "Done!"
