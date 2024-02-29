run: 
	go run cmd/auth/main.go

build:
	go build cmd/auth/main.go

test: 
	go test -cover ./...
