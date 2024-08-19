build:
	@go build -o bin/fs
	@echo "Built binary successfully!"

run: build
	@./bin/fs

test:
	@go test ./...
