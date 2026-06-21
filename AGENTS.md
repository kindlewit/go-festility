# Festility — Agent Guide

## Project Overview

**Festility** (`github.com/kindlewit/go-festility`) is a Go-based REST API for managing film festivals, cinemas, movies, and schedules. It uses the **Gin** web framework and **MongoDB** as its database.

The project also includes a **React/Next.js** front-end application located in the `web/` folder.

---

## Stack

| Layer      | Technology                          |
|------------|-------------------------------------|
| Language   | Go 1.22+                            |
| API        | Gin (`github.com/gin-gonic/gin`)    |
| Database   | MongoDB (`go.mongodb.org/mongo-driver`) |
| Testing    | `github.com/stretchr/testify`       |
| Frontend   | Next.js / React (TypeScript) — `web/` |
| Container  | Docker / Docker Compose             |

---

## Running the Project

### Locally
```bash
go build          # builds binary named after the module (go-festility)
go run .          # starts the server
```

### Docker
```bash
# Build and run
docker build -t festility .
docker run --rm -it -p 8080:8080 festility

# Or with Compose
docker-compose up --build
```

### MongoDB (standalone)
```bash
docker run -d -p 27017:27017 mongo
docker exec -it <container_name> mongosh
```

---

## Frontend (`web/`)

The React/Next.js application lives entirely under `web/`. It is a separate application from the Go backend.

```
web/
  pages/          # Next.js pages (routes)
  components/     # Reusable React components
  entities/       # TypeScript type definitions
  util/           # Frontend utility functions
  styles/         # CSS / style modules
  constants.ts    # Shared frontend constants
  data.ts         # Static or seed data
```

---

## Project Structure Conventions (Domain-Driven Design)

This project follows a **Domain-Driven Design (DDD)** structure for the Go backend.

### Target layout

```
src/
  db/             # Database connection and access module
  api/            # Gin router and all middleware
  constants/      # Shared constants (*.go)
  <domain>/       # One folder per domain (e.g. cinema, festival, movie, schedule)
    handler.go    # HTTP handlers for the domain
    service.go    # Business logic
    repository.go # Database access layer
    models.go     # Domain models / structs
web/              # React / Next.js frontend application
main.go           # Entry point to start the application
```

### Rules

- **`src/api/`** — contains only the Gin router setup and middleware (auth, logging, CORS, etc.). No business logic here.
- **`src/<domain>/`** — each domain (e.g. `cinema`, `festival`, `movie`, `schedule`) gets its own folder under `src/`. The four standard files are:
  - `handler.go` — binds HTTP routes to service calls; handles request parsing and response formatting.
  - `service.go` — implements business rules; calls the repository.
  - `repository.go` — all direct database interactions (MongoDB queries).
  - `models.go` — struct definitions and any domain-specific types.
- **`src/constants/`** — shared, cross-domain constants (`*.go`). Domain-specific constants belong inside the domain folder.
- Do **not** place business logic in handlers or database calls in services — keep the layers strictly separated.

---

## Testing

This project follows a **Test-Driven Development (TDD)** approach aligned with the DDD layer structure. Each layer is tested in isolation using fakes — never real infrastructure.

### Core principle: test each layer with a fake of the layer below

| Layer under test | Fake used           | What is verified                          |
|------------------|---------------------|-------------------------------------------|
| `repository.go`  | `FakeDatabase`      | Correct queries & data mapping            |
| `service.go`     | `FakeRepository`    | Business rules, error handling            |
| `handler.go`     | `FakeService`       | HTTP status codes, request/response shape |

### Test folder structure

Tests for `src/` code live **inside each domain folder**, co-located with the production code. The Go convention of `_test.go` suffix is used throughout.

```
src/
  <domain>/
    handler.go
    handler_test.go       # tests handlers using FakeService
    service.go
    service_test.go       # tests services using FakeRepository
    repository.go
    repository_test.go    # tests repositories using FakeDatabase
    models.go
    fakes_test.go         # FakeService, FakeRepository, FakeDatabase definitions
                          # (unexported, only available within the domain's test files)
```

Example for the `cinema` domain:

```
src/cinema/
  models.go
  handler.go
  handler_test.go
  service.go
  service_test.go
  repository.go
  repository_test.go
  fakes_test.go
```

### Fake conventions

- All fakes are defined in `fakes_test.go` within the domain package.
- Fakes implement the same interface as the real component they replace.
- This requires each layer to depend on an **interface**, not a concrete type:
  - `repository.go` accepts a `DatabaseClient` interface (satisfied by `src/db.Database` in production, `FakeDatabase` in tests).
  - `service.go` accepts a `Repository` interface (satisfied by the real repository in production, `FakeRepository` in tests).
  - `handler.go` accepts a `Service` interface (satisfied by the real service in production, `FakeService` in tests).
- Fakes record calls and return pre-configured responses — no real MongoDB or HTTP calls in unit tests.
- Handler tests use `net/http/httptest` to fire requests against a `gin.Engine` wired with the `FakeService`.

### Running tests

```bash
go test ./...                        # all tests
go test ./src/cinema/...             # one domain
go test ./src/cinema/... -v          # verbose output
go test ./src/cinema/... -run Name   # single test by name
```

### Legacy tests

The existing tests in `tests/` predate this structure and may use real infrastructure or older patterns. New tests must follow the TDD-DDD approach described above.
