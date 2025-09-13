.PHONY: test

run:
	@echo "Running the application..."
	@go run ./app/main.go

build:
	@ rm -f rest-api
	@echo "Building the application..."
	@go build -o rest-api ./app/main.go
	@echo "Done"

test:
	@go test -cover -coverprofile=coverage.out ./test/... -coverpkg=./app/...

test-v:
	@go test -v -cover -coverprofile=coverage.out ./test/... -coverpkg=./app/...

test-coverage:
	@go tool cover -html=coverage.out

docker-build:
	@docker build -f docker/Dockerfile -t jmticonap/rest-api-boilerplate .

docker-up:
	@docker compose -f docker/docker-compose.yml up --build --abort-on-container-exit