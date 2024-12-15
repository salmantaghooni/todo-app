.PHONY: run test benchmark

# Run the application using Docker Compose
run:
	docker-compose up --build

# Run all unit tests
test:
	go test ./... -cover

# Run all benchmarks
benchmark:
	go test -bench=.