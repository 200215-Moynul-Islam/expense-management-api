# Expense Management API

A RESTful expense tracking API built with **Go** and the **Beego v2** framework. Users can register, authenticate, organize expenses into categories, and manage their personal spending records.

## Features

- JWT-based authentication (register / login)
- Category management (CRUD, scoped per user)
- Expense management (CRUD) with filtering, pagination, and sorting
- User profile retrieval and update
- Layered architecture: controllers → services → repositories → models
- PostgreSQL persistence via Beego ORM
- Dockerized PostgreSQL for local development

## Tech Stack

- **Language:** Go 1.26
- **Framework:** [Beego v2](https://github.com/beego/beego)
- **Database:** PostgreSQL 17
- **Auth:** JWT ([golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt))

## Project Structure

```
.
├── conf/               # App configuration (app.conf)
├── controllers/        # HTTP handlers (auth, category, expense, user, health)
├── database/           # DB connection/initialization
├── dto/                # Request/response data transfer objects
├── errors/             # Centralized sentinel errors
├── middlewares/        # Auth middleware (JWT filter)
├── migrations/         # SQL schema migrations
├── models/             # ORM models (User, Category, Expense)
├── postman/            # Postman collection for API testing
├── repositories/       # Data access layer
├── routers/            # Route definitions
├── services/           # Business logic layer
├── utils/              # JWT, response helpers, validators
└── main.go
```

## Getting Started

### Prerequisites

- Go 1.26+
- Docker & Docker Compose (for PostgreSQL)
- [Bee CLI](https://github.com/beego/bee) — Beego's dev tool (`go install github.com/beego/bee/v2@latest`)

### 1. Clone the repository

```bash
git clone https://github.com/200215-Moynul-Islam/expense-management-api.git
cd expense-management-api
```

### 2. Create your local config

Copy the sample config and adjust values as needed:

```bash
cp conf/app.conf.sample conf/app.conf
```

`conf/app.conf.sample` ships with these defaults:

```ini
appname = expense-management-api
httpport = 8080
runmode = dev
autorender = false
copyrequestbody = true
EnableDocs = true

# Database configuration(PostgreSQL)
POSTGRES_DB=expense_management_db
POSTGRES_USER=expense_management_db_user
POSTGRES_PASSWORD=your_password
POSTGRES_HOST=localhost
POSTGRES_PORT=5433
POSTGRES_SSLMODE=disable

# JWT Configuration
JWT_SECRET = your_secret_key
JWT_EXPIRATION_MINUTES=30
```

Set a real `JWT_SECRET` and `POSTGRES_PASSWORD`. `POSTGRES_HOST` is already set to `localhost` since only PostgreSQL runs in Docker — the app itself runs directly on your host machine.

**If you change any value from the sample**, keep these in sync elsewhere:

| If you change...                                    | Also update...                                                                                                                                         |
| --------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `httpport`                                          | The port you use to call the API (e.g. in Postman's base URL variable)                                                                                 |
| `POSTGRES_DB`, `POSTGRES_USER`, `POSTGRES_PASSWORD` | The same values in `docker-compose.yml`, and the connection string used for migrations (step 4 below)                                                  |
| `POSTGRES_HOST`, `POSTGRES_PORT`                    | The connection string used for migrations (step 4 below); `POSTGRES_PORT` must also match the host port mapped in `docker-compose.yml` (`"5433:5432"`) |
| `JWT_SECRET`                                        | Nothing else — it only needs to be consistent across app restarts so existing tokens stay valid                                                        |

### 3. Start PostgreSQL

```bash
docker-compose up -d
```

This starts PostgreSQL 17 in Docker, exposing container port `5432` on host port `5433` (per `docker-compose.yml`). This host port must match `POSTGRES_PORT` in `conf/app.conf`.

### 4. Run database migrations

Migrations live as plain SQL files in `migrations/` and are applied with [golang-migrate](https://github.com/golang-migrate/migrate) (install it separately; it isn't a Go module dependency of this project).

Install the CLI:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Apply all pending migrations, using the same DB name/user/password/host/port you set in `conf/app.conf`:

```bash
migrate -path migrations \
  -database "postgres://expense_management_db_user:your_password@localhost:5433/expense_management_db?sslmode=disable" \
  up
```

To roll back the last migration:

```bash
migrate -path migrations \
  -database "postgres://expense_management_db_user:your_password@localhost:5433/expense_management_db?sslmode=disable" \
  down 1
```

### 5. Run the application with Bee

```bash
bee run
```

Bee compiles and runs the app, and live-reloads on file changes. The API will be available at `http://localhost:8080` (the `httpport` set in `conf/app.conf`).

## Configuration

| Key                                                   | Description                 |
| ----------------------------------------------------- | --------------------------- |
| `httpport`                                            | Port the server listens on  |
| `runmode`                                             | `dev` or `prod`             |
| `POSTGRES_DB` / `POSTGRES_USER` / `POSTGRES_PASSWORD` | Database credentials        |
| `POSTGRES_HOST` / `POSTGRES_PORT`                     | Database host/port          |
| `POSTGRES_SSLMODE`                                    | Database SSL mode           |
| `JWT_SECRET`                                          | Secret used to sign JWTs    |
| `JWT_EXPIRATION_MINUTES`                              | Token expiration in minutes |

## API Reference

All endpoints are prefixed with `/api/v1`. Endpoints marked 🔒 require an `Authorization: Bearer <token>` header.

### Health

| Method | Endpoint  | Description          |
| ------ | --------- | -------------------- |
| GET    | `/health` | Service health check |

### Auth

| Method | Endpoint         | Description              |
| ------ | ---------------- | ------------------------ |
| POST   | `/auth/register` | Register a new user      |
| POST   | `/auth/login`    | Log in and receive a JWT |

### Users 🔒

| Method | Endpoint    | Description                 |
| ------ | ----------- | --------------------------- |
| GET    | `/users/me` | Get current user profile    |
| PUT    | `/users/me` | Update current user profile |

### Categories 🔒

| Method | Endpoint          | Description                          |
| ------ | ----------------- | ------------------------------------ |
| POST   | `/categories`     | Create a category                    |
| GET    | `/categories`     | List categories for the current user |
| GET    | `/categories/:id` | Get a category by ID                 |
| PUT    | `/categories/:id` | Update a category                    |
| DELETE | `/categories/:id` | Delete a category                    |

### Expenses 🔒

| Method | Endpoint        | Description                                   |
| ------ | --------------- | --------------------------------------------- |
| POST   | `/expenses`     | Create an expense                             |
| GET    | `/expenses`     | List expenses (supports filtering/pagination) |
| GET    | `/expenses/:id` | Get an expense by ID                          |
| PUT    | `/expenses/:id` | Update an expense                             |
| DELETE | `/expenses/:id` | Delete an expense                             |

**Expense query parameters:** `category_id`, `from_date`, `to_date` (`YYYY-MM-DD`), `page`, `limit`, `sort_by` (`created_at`, `expense_date`, `amount`), `sort_order` (`asc`, `desc`).

## Postman Collection

A ready-to-use Postman collection is included at `postman/Expense Management API.postman_collection.json`.

1. Open Postman and click **Import**.
2. Select `postman/Expense Management API.postman_collection.json` from this repo.
3. Update the collection's base URL variable to match your local server (default `http://localhost:8080`).
4. Run **Auth → Login** (or **Register** then **Login**) first. The collection's `access_token` variable is set automatically from the response, so all subsequent 🔒 requests are authenticated without any manual copying.

## Database Schema

- **users** — `id`, `name`, `email` (unique), `password`, timestamps
- **categories** — `id`, `user_id` (FK), `name` (unique per user), timestamps
- **expenses** — `id`, `user_id` (FK), `category_id` (FK), `title`, `amount`, `note`, `expense_date`, timestamps

See `migrations/000001_create_initial_schema.up.sql` for full DDL.

## Response Format

All API responses follow a consistent envelope:

```json
{
  "success": true,
  "message": "Descriptive message.",
  "data": {}
}
```
