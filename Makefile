# Makefile for coco-sml semantic search API server

.PHONY: run

run:
	go run main.go config.json

# Optional: build target
.PHONY: build
build:
	go build -o coco-sml main.go

# Optional: clean target
.PHONY: clean
clean:
	rm -f coco-sml
	rm -f ./data/server.pid

