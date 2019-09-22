## Calendar server

### Menu

- [Get a Repo](#get-a-repo)
- [Build](#build)
- [Run](#run)
- [Run With Docker](#run-with-docker)
- [Configuration Example](#configuration-example)

### Get a Repo
```bash
>>> git clone https://github.com/antik9/microservice-go.git
>>> cd microservice-go
```

### Build

```bash
>>> make
```

### Run

```
>>> calendar-migrate # run migrations on postgres database
>>> calendar-server # run grpc server
>>> calendar-client # create your own events
>>> calendar-enqueue # enqueue coming events to rabbitmq
>>> calendar-dequeue # dequeue coming events from rabbitmq and print to console
```

### Run With Docker
```
>>> cd docker && docker-compose up
```

### Configuration Example

```yaml
database:
    name: calendar
    username: calendar
    host: 127.0.0.1
    password: calendar
    backend: postgres

rabbit:
   username: guest
   password: guest
   host: localhost
   port: 5672
```
