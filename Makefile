.PHONY: build

build:
	go build -o out/brisk main.go

release:
	go build -o release/macos/brisk main.go