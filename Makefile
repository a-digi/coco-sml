# Makefile for coco-sml semantic search API server

.PHONY: run
run: build
	nohup ./app/coco-sml start --data-dir=./data --config=config.json > server.log 2>&1 &
	@echo "Server started in background. Logs: server.log"

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

.PHONY: build-frontend
build-frontend:
	cd src/front/app && npm install && npm run build

.PHONY: serve-frontend
serve-frontend:
	cd src/front/app && npm run start
