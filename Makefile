.PHONY: build

build:
	./build.sh

release:
	go build -o release/macos/brisk main.go