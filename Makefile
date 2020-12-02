.PHONY: build run clean test
build:
	go build -o build/kitchendb pkg/kitchen/main.go
run:
	./build/kitchendb
clean:
	rm -r build/*
test:
	go test ./...