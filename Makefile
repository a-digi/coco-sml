# Makefile for coco-sml semantic search API server

.PHONY: run
run: build
	./app/coco-sml --data-dir=./data --config=config.json

.PHONY: run-dev
run-dev:
	go run main.go config.json

.PHONY: stop-dev
stop-dev:
	go run stop_main.go

.PHONY: build
build:
	mkdir -p app
	go build -o app/coco-sml main.go

.PHONY: clean
clean:
	rm -f app/coco-sml
	rm -f ./data/server.pid

.PHONY: stop
stop:
	./app/coco-sml stop
