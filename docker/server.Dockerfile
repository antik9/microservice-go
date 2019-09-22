FROM golang:1.13.0-buster

WORKDIR /opt/microservice-go
RUN git clone https://github.com/antik9/microservice-go /opt/microservice-go
RUN go build -o calendar-server cmd/server/main.go
RUN cp /opt/microservice-go/docker/conf.yaml /opt/microservice-go/conf.yaml

ENTRYPOINT ["./calendar-server"]
