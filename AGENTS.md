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

Tests live in `tests/`. Run them with:
```bash
go test ./...
```
