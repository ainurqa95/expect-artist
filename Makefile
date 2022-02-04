.PHONY:
.SILENT:

build: 
	go build -o ./main cmd/bot/main.go

run: build
	./.main

docker-build:
	docker-compose build expect-artist

up-container:
	docker-compose up expect-artist
