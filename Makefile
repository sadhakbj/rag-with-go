run:
	go run cmd/main.go

build:
	mkdir -p bin
	go build -o bin/app cmd/main.go
