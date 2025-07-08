# crud_postgres

A minimal CRUD API built with Go and PostgreSQL using `database/sql`, `lib/pq`, and `chi` router.

## 🚀 Quick Start

### Setup PostgreSQL

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

### Environment

Create a .env file:

```env
DATABASE_URL=postgres://username:password@localhost/users_db?sslmode=disable
```

### Install & Run

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

### Output

### 🔹 Create User

**Endpoint:** `POST /user`  
**Description:** Creates a new user in the PostgreSQL database.

#### 📥 Request Body (JSON)

```json
{
  "name": "Dhanush",
  "email": "dhanush@example.com"
}
```

![Screenshot 2025-07-08 at 4 42 46 PM](https://github.com/user-attachments/assets/6b404628-9b29-4a83-9a68-a71e4160ef57)

### 🔹 Get User by ID

**Endpoint:** `GET /user/:id`  
**Description:** Retrieves a user by their unique ID from the PostgreSQL database.

Example:  
`GET /user/1`

#### 📤 Response (200 OK)

```json
{
  "id": 1,
  "name": "Dhanush",
  "email": "dhanush@example.com"
}
```

![Screenshot 2025-07-08 at 4 43 22 PM](https://github.com/user-attachments/assets/e3c89ed6-9e8b-47b4-8896-c9b0a461561b)

### 🔹 Update User

**Endpoint:** `PUT /user/:id`  
**Description:** Updates an existing user's name and email.

#### 📥 Request Body (JSON)

```json
{
  "name": "Updated Name",
  "email": "updated@example.com"
}
```

### 🔹 Delete User

**Endpoint:** `DELETE /user/:id`  
**Description:** Deletes a user from the PostgreSQL database by their unique ID.

Example:  
`DELETE /user/1`

