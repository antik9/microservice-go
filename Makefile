.PHONY: install

install: migrate server client

migrate:
	go build -o calendar-migrate cmd/migration/main.go

server:
	go build -o calendar-server cmd/server/main.go

client:
	go build -o calendar-client cmd/client/main.go
