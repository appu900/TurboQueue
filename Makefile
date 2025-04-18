.PHONY: build run test migrate clean

# Build the Go binary into ./bin
build:
	go build -o bin/turboQueue ./cmd/turboQueue

# Run the built binary from project root (so ./config is found)
run: build
	./bin/turboQueue

# Run tests
test:
	go test ./...

# Placeholder for DB migrations (adjust command as needed)
migrate:
	# Example: goose -dir migrations postgres "user=postgres password=password dbname=qstash sslmode=disable" up

# Clean the bin directory
clean:
	rm -rf bin
