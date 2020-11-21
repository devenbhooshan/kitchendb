.PHONY: build
build:
	go build -o build/main src/main.go
run:
	./build/main