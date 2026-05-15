.PHONY: test fmt lint

test:
	go test ./...

fmt:
	gofmt -w .

lint:
	golangci-lint run ./...
