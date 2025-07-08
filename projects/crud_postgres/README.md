# crud_postgres

A minimal CRUD API built with Go and PostgreSQL using `database/sql`, `lib/pq`, and `chi` router.

## 🚀 Quick Start

### 1. Setup PostgreSQL

```sql
CREATE DATABASE users_db;
CREATE USER username WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE users_db TO username;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO username;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO username;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);
```

### 2. Environment

Create a .env file:

```env
DATABASE_URL=postgres://username:password@localhost/users_db?sslmode=disable
```

### 3. Install & Run

```
go mod init crud_postgres

go get github.com/lib/pq
go get github.com/go-chi/chi/v5
go get github.com/joho/godotenv
go get github.com/go-chi/chi/v5/middleware

go mod tidy
go run main.go
```

Server runs at http://localhost:8080