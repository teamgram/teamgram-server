all: build run

build:
	docker build -t chatengine/server:latest .

run:
	docker-compose up -d