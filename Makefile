all: build

build: build-client build-server

build-server:
	go build -o bin/queue-server ./cmd/server

build-client:
	go build -o bin/queue-cli ./cmd/client