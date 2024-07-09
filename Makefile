dev:
	go run cmd/server.go

test:
	go test -v ./...

test-cover:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test-it:
	go test -v -tags=integration  ./... 
