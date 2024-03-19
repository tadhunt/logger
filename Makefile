export GOPRIVATE=github.com/tadhunt

all:
	go mod tidy
	go generate
	go vet
	staticcheck
	go build

test: all
	go test -v ./...
