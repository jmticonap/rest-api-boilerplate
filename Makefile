.PHONY: test

run:
	@echo "Running the application..."
	@go run ./app/cmd/main.go

build:
	@ rm -f router-schema
	@echo "Building the application..."
	@go build -o router-schema ./app/cmd/main.go
	@echo "Done"

test:
	@go test -v -cover -coverprofile=coverage.out ./... -coverpkg=./app/...

test-coverage:
	@go tool cover -html=coverage.out
