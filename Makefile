export GOPRIVATE=github.com/tadhunt

all:
	go mod tidy
	go generate
	go vet
	staticcheck
	go build
	GOOS=windows GOARCH=amd64 go build

test: all
	go test -v ./...
