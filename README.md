## Calendar server

### Get a Repo
```bash
>>> go get -u github.com/antik9/microservice-go
```

### Build

```bash
>>> make
```

### Uninstall
```bash
>>> make clean
```

### Run

```
>>> calendar-migrate # run migrations on postgres database
>>> calendar-server -db postgres # or -db memory
>>> calendar-client
```

### Configuration Example

```yaml
database:
    name: calendar
    username: calendar
    host: 127.0.0.1
    password: calendar
```
