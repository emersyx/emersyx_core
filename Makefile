.PHONY: emcore test

emcore:
	@echo "Building emersyx core..."
	-@rm emersyx
	@dep ensure
	@go build -o emersyx ./emcore/*

test:
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l emcore/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
