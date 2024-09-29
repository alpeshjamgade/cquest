APP_BINARY=cquestApp
DATABASE_URL="postgresql://postgres:postgres@localhost:5432/cquest?sslmode=disable"
MIGRATION_PATH="./internal/migrations/"

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

migration_create:
	migrate create -ext sql -dir ${MIGRATION_PATH} -seq ${name}

migration_up:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URL} -verbose up

migration_down:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URL} -verbose down

migration_fix:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URL} force VERSION
