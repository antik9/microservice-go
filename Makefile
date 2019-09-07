.PHONY: install

install: migrate server client

migrate:
	go build -o calendar-migrate cmd/migration/main.go
	cp calendar-migrate ${GOPATH}/bin

server:
	go build -o calendar-server cmd/server/main.go
	cp calendar-server ${GOPATH}/bin

client:
	go build -o calendar-client cmd/client/main.go
	cp calendar-client ${GOPATH}/bin

clean:
	rm ${GOPATH}/bin/calendar-migrate ${GOPATH}/bin/calendar-server ${GOPATH}/bin/calendar-client
