run: build
	@./bin/api

 build:
	@go build -o bin/api ./cmd/task-manager/

 test:
	@go test -v ./internal/tests/...