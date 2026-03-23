# Dashboard API BFF

A Backend for Frontend (BFF) REST API built in Go that aggregates user profile and todo data
from [DummyJSON](https://dummyjson.com) into a single dashboard endpoint.
Built as a technical assessment.

| Package    | File               | Purpose                                            |
|------------|--------------------|----------------------------------------------------|
| `main`     | `config.go`        | Loads configuration from environment variables.    |
|            | `main.go`          | Wires dependencies and starts the server.          |
| `domain`   | `user.go`          | User domain logic/rules.                           |
|            | `user_test.go`     | Unit tests for User.                               |
|            | `todo.go`          | Todo domain logic/rules.                           |
|            | `todo_test.go`     | Unit tests for Todo.                               |
| `response` | `response.go`      | Aggregates user and todos into dashboard response. |
|            | `response_test.go` | Unit tests for dashboard aggregation logic.        |
| `internal` | `client.go`        | HTTP client for DummyJSON API.                     |
|            | `service.go`       | Fetches user and todos concurrently.               |
|            | `handler.go`       | Handles HTTP request and writes JSON response.     |

## How to run the application

```shell
# run with default config
go run ./cmd/api

# build and run (osx|linux)
go build -o bff-dashboard-api ./cmd/api && ./bff-dashboard-api

# build and run (powershell)
go build -o bff-dashboard-api.exe ./cmd/api; ./bff-dashboard-api.exe

# run with custom config (osx|linux)
DUMMYJSON_BASE_URL=https://dummyjson.com DASHBOARD_TIMEOUT=2s go run ./cmd/api

# run with custom config (powershell)
$env:DUMMYJSON_BASE_URL="https://dummyjson.com"; $env:DASHBOARD_TIMEOUT="2s"; go run ./cmd/api
```

## How to test the application

```shell
curl http://localhost:8080/dashboard/8
```

**Sample response with all data available:**

```json
{
  "id": 8,
  "full_name": "Ava Taylor",
  "status": "Rookie",
  "pending_task_count": 1,
  "next_urgent_task": "Take a scenic horseback riding tour",
  "error_warning": null
}
```

**Sample response with partial failure handling simulation:**

```json
{
  "id": 8,
  "full_name": "Ava Taylor",
  "status": "Rookie",
  "pending_task_count": 0,
  "next_urgent_task": null,
  "error_warning": "Todos Unavailable"
}
```

> The simulated delay is controlled by the current second at execution time.
> Todos endpoint adds a 3000ms delay when `second < 30 && second % 3 == 0`,
> triggering the 2s timeout and returning `"Todos Unavailable"`.

## How to run the tests

```shell
# run all tests
go test ./...

# run all tests with verbose output
go test -v ./...
```

## External Data Sources

1. User Profile: https://dummyjson.com/users/:id
2. User Todos: https://dummyjson.com/todos/user/:id

<!--
## Checklist

- [x] Base server with GET `dashboard/:id` endpoint.
- [x] Client, external data source: `/users/:id`
- [x] Client, external data source: `/todos/user/:id`
- [x] Client, strict 2 seconds timeout.
- [x] Test client `user`
- [x] Test client `todo`
- [x] **Fetch calls in parallel.**
- [x] Business rule, full name (first + last).
- [x] Business rule, age > 50 with status Veteran, otherwise Rookie.
- [x] Business rule, count of pending tasks.
- [x] Business rule, title of the first pending task.
- [x] Business rule, filter todos completed.
- [x] Business rule, partial failure handling with warning.
- [x] Test and validate business rules (QA).
- [x] Unit tests.
- [x] README.
- [x] GIT.

### Expected output

```json
{
  "id": 1,
  "full_name": "Terry Medhurst",
  "status": "Rookie",
  "pending_task_count": 2,
  "next_urgent_task": "Do something important",
  "error_warning": null
}
```
-->