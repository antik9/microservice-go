## Calendar server

### Get a Repo
```bash
>>> go get -u github.com/antik9/microservice-go
```

### Run

```bash
>>> microservice-go -mode server -db postgres # or -db memory

>>> microservice-go -mode client
```

### Configuration Example

```yaml
database:
    name: calendar
    username: calendar
    host: 127.0.0.1
    password: calendar
```
