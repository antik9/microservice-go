.PHONY: install

install: migrate server client enqueue dequeue

migrate:
	go build -o calendar-migrate cmd/migration/main.go

server:
	go build -o calendar-server cmd/server/main.go

client:
	go build -o calendar-client cmd/client/main.go

enqueue:
	go build -o calendar-enqueue cmd/enqueue/main.go

dequeue:
	go build -o calendar-dequeue cmd/dequeue/main.go
