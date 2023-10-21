build:
	@go build - bin/goapi

run: build
	@./bin/goapi

test: 
	@go test -v ./...