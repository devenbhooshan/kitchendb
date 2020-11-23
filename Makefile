.PHONY: build run clean
build:
	go build -o build/kitchendb pkg/kitchen/main.go
run:
	./build/kitchendb

clean:
	rm -r build/*