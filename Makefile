emersyx: goget
	@go build -o emersyx ./emcore/*

.PHONY: goget
goget:
	go get emersyx.net/emersyx_apis/emcomapi
	go get emersyx.net/emersyx_apis/emircapi
	go get emersyx.net/emersyx_apis/emtgapi
	go get emersyx.net/emersyx_log/emlog
	go get github.com/BurntSushi/toml
	go get github.com/golang/lint/golint

.PHONY: test
test: emersyx
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l emcore/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
	@cd emcore; go test -v -conffile ../config.toml
