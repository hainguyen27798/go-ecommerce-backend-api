# app name
APP_NAME = server

run:
	go run ./cmd/${APP_NAME}/

build:
	go build -o backend ./cmd/${APP_NAME}