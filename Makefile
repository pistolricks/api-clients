## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api

## client/install: install client dependencies
.PHONY: client/install
client/install:
	cd client && npm install

## client/dev: run client in development mode
.PHONY: client/dev
client/dev:
	cd client && npm run dev

## client/build: build client for production
.PHONY: client/build
client/build:
	cd client && npm run build

## run/all: run both API and client
.PHONY: run/all
run/all:
	go run ./cmd/api & cd client && npm run dev
