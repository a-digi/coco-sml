# Makefile for coco-sml semantic search API server

.PHONY: run

run:
	go run main.go config.json

.PHONY: stop

stop:
	go run -e 'package main; import "github.com/a-digi/coco-sml/src/server"; func main() { server.StopServer("./data") }'

# Optional: build target
.PHONY: build
build:
	go build -o coco-sml main.go

# Optional: clean target
.PHONY: clean
clean:
	rm -f coco-sml
	rm -f ./data/server.pid
