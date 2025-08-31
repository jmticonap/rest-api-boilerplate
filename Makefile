.PHONY: test

test:
	@go test -v -cover -coverprofile=coverage.out ./... -coverpkg=./lib/...
	@go tool cover -html=coverage.out
