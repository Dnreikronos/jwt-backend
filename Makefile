build:
	@go build -o bin/jwt-backend cmd/main.go

run: build
	@./bin/jwt-backend cmd/main.go

test:
	@go test -v ./...
