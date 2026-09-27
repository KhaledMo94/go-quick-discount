# Quick Discount API

A Go REST API for the Quick Discount platform, built with [Gin](https://github.com/gin-gonic/gin) and [GORM](https://gorm.io). It shares its MySQL schema with an existing Laravel application (Sanctum-compatible personal access tokens, Laravel-style polymorphic types).

## Tech stack

- **Go** 1.26+
- **Gin** — HTTP router
- **GORM** + **MySQL** — persistence
- **Redis** — cache
- **Meilisearch** — search
- **lumberjack** — rotating log files
- **godotenv** — `.env` loading

## Project structure

```text
cmd/myapp/            Application entry point (composition root)
internal/
  config/             Environment configuration
  db/                 MySQL and GORM connections
  redis/              Redis client
  meliesearch/        Meilisearch client
  logger/             Structured logging (slog + rotating files)
  i18n/               Translation loader
  models/             GORM models
  types/              Shared custom types and query scopes
  repository/         Data access (GORM implementations)
  service/            Business logic
    interfaces/       Repository abstractions used by services
  handler/            HTTP handlers, routing, middleware
  helper/             Stateless helpers (token hashing, parsing)
  server/             HTTP server lifecycle
lang/                 Translation files (e.g. ar.json)
logs/                 Runtime logs (git-ignored)
```

## Architecture

Requests flow through layered packages, wired together in `cmd/myapp/main.go`:

```text
handler  →  service  →  interfaces  ←  repository  →  database
```

- **Services depend on interfaces**, never on concrete repositories or GORM.
- **Repositories** only handle persistence; crypto and parsing live in `internal/helper`.
- **Interfaces are small and role-based** where consumers differ (e.g. `TokenIssuer`, `TokenResolver`, `TokenRevoker` for personal access tokens).
- Cross-cutting concerns such as caching should be added as decorators implementing the same repository interface, so services stay unchanged.

## Getting started

### Prerequisites

- Go 1.26 or newer
- MySQL
- Redis
- Meilisearch

### Configuration

Create a `.env` file in the project root (it is git-ignored). If it is missing, system environment variables are used.

```dotenv
APP_URL=http://localhost:8000
APP_LOCALE=en
APP_FALLBACK_LOCALE=en

DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=quick_discount
DB_USERNAME=root
DB_PASSWORD=

REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

MEILISEARCH_HOST=http://127.0.0.1:7700
MEILISEARCH_KEY=
SCOUT_PREFIX=
```

`DB_HOST`, `DB_PORT`, `DB_DATABASE`, and `DB_USERNAME` are required.

### Run

```bash
go mod download
go run ./cmd/myapp
```

The server listens on `:8000` and shuts down gracefully on `SIGINT` / `SIGTERM`.

### Build

```bash
go build -o bin/quick-discount ./cmd/myapp
```

## API

| Method | Path                   | Description                                   |
|--------|------------------------|-----------------------------------------------|
| GET    | `/health`              | Health of MySQL, Redis, and Meilisearch       |
| GET    | `/api/main-categories` | List active main categories                   |

`/health` returns `200 {"status":"ok"}` when all dependencies respond, otherwise `503 {"status":"degraded"}`.

### Response format

```json
{
  "status": "success",
  "message": "ok",
  "data": []
}
```

### Localization

The response language is resolved from the `lang` query parameter, then the `Accept-Language` header, then `APP_LOCALE`. Translatable fields stored as JSON (e.g. `{"en": "...", "ar": "..."}`) are returned in the resolved locale.

```bash
curl "http://localhost:8000/api/main-categories?lang=ar"
```

## Authentication

Personal access tokens follow the Laravel Sanctum format `"{id}|{secret}"`. Only the SHA-256 hash of the secret is stored, and verification uses a constant-time comparison, so tokens issued by the Laravel app remain valid here.

## Logging

Logs are written with `log/slog` to `logs/app.log` with automatic rotation.
