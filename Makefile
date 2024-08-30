# app name
APP_NAME = server

run:
	MODE=dev go run ./cmd/${APP_NAME}/

build:
	go build -o backend ./cmd/${APP_NAME}

wire:
	cd internal/wires && wire