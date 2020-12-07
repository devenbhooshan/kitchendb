.PHONY: build run clean test
build:
	go build -o build/kitchendb pkg/kitchen/main.go
run:
	./build/kitchendb
clean:
	go clean -testcache && rm -rf build/*
test:
	go test ./...
deps:
	go get -v -t -d ./...