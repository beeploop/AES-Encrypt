build:
	@go build -o bin/aes-encrypt main.go

run:
	@go run main.go

clean:
	@rm -rf bin

.PHONY: build run clean
