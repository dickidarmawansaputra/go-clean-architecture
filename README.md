# Go Clean Architecture With Fiber
<p>This repository is Golang implementation of Clean Code Architecture using Fiber framework.</p>

![Go](go.jpeg)

## Getting Started
### Start application using docker for production
1. Clone repository
```bash
git clone https://github.com/dickidarmawansaputra/go-clean-architecture.git
```
2. Running docker compose & make sure you have installed docker in your machine
```bash
docker compose up -d
```
3. Migrate database
```bash
migrate -path database/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up
```
### Start application using docker for development
1. Using go live reload & add to `Dockerfile`
```dockerfile
FROM golang:alpine3.21

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

CMD ["air", "-c", ".air.toml"]
```
2. Initialize the `.air.toml` configuration file
```bash
air init
```
3. Change build commad to `cmd/web` directory in `.air.toml` file
```toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ./cmd/web"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  post_cmd = []
  pre_cmd = []
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_error = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  silent = false
  time = false

[misc]
  clean_on_exit = false

[proxy]
  app_port = 0
  enabled = false
  proxy_port = 0

[screen]
  clear_on_rebuild = false
  keep_scroll = true

```

## HTTP Methods for RESTful APIs

| HTTP Method | Description                         |
| ----------- | ----------------------------------- |
| GET         | Retrieve a resource.                |
| POST        | Create a new resource.              |
| PUT         | Update a resource (full update).    |
| PATCH       | Update a resource (partial update). |
| DELETE      | Remove a resource.                  |

## Error Handling and Standard HTTP Status Codes

| Status Code               | Description                                       |
| ------------------------- | ------------------------------------------------- |
| 200 OK                    | Successfully returned data.                       |
| 201 Created               | Successfully created a resource.                  |
| 204 No Content            | Successfully deleted a resource.                  |
| 400 Bad Request           | Invalid request (e.g., malformed JSON).           |
| 401 Unauthorized          | Authentication required.                          |
| 403 Forbidden             | Access denied.                                    |
| 404 Not Found             | Resource not found.                               |
| 409 Conflict              | Conflict with the current state of a resource.    |
| 422 Unprocessable Entity  | Validation or processing error.                   |
| 500 Internal Server Error | Server encountered an error.                      |
| 503 Service Unavailable   | Server temporarily unavailable.                   |

## Tech Stack & Tools

- Framework: [Fiber](https://gofiber.io)
- Database: [PostgreSQL](https://github.com/go-gorm/postgres)
- ORM: [GORM](https://gorm.io)
- Config: [Viper](https://github.com/spf13/viper)
- Validator: [Go Validator](https://github.com/go-playground/validator)
- Log: [Logrus](https://github.com/sirupsen/logrus)
- APIs Docs: [Swagger](https://github.com/gofiber/swagger)
- Container: [Docker](https://www.docker.com)

## Author

[Dicki D. Saputra](http://github.com/dickidarmawansaputra)
