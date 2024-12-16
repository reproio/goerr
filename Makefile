.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: fmt
fmt:
	go mod tidy
	golangci-lint run --fix
	go fmt ./...

.PHONY: install-tools
install-tools:
	@cat tools.go | awk -F'"' '/_/ {print $$2}' | xargs -tI {} go install {}
