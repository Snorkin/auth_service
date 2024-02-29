run: 
	go run cmd/auth/main.go

build:
	go build -o bin/ cmd/auth/main.go

test: 
	go test -cover ./...
