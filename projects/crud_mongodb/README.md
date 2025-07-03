# crud_mongodb

A simple RESTful CRUD API built in Go that connects to MongoDB using the official Go MongoDB driver. This project demonstrates how to perform Create, Read, and Delete operations on user data, with basic validation and clean project structure.

## How to run
```bash
cd crud_mongodb
```
Install dependencies
```
go mod tidy
```
Make sure your mongodb in local is up and running
Run the application
```
go run main.go
```

### 🔹 Get User by ID

**Endpoint:** `GET /user/:id`  
**Description:** Retrieves a single user document from MongoDB by its unique ObjectID.

Example:  
`GET /user/60c8d0f2e3b5c47a2c84f6d9`

#### 📤 Successful Response (200 OK)

```json
{
  "id": "60c8d0f2e3b5c47a2c84f6d9",
  "name": "Dha",
  "email": "dha@example.com"
}
```

---

### 🔹 Create User

**Endpoint:** `POST /user`  
**Description:** Creates a new user in the MongoDB collection.


#### 📥 Request Body (JSON)

```json
{
  "name": "Dhanush",
  "email": "dhanush@example.com"
}
```

---

### 🔹 Delete User

**Endpoint:** `DELETE /user/:id`  
**Description:** Deletes a user from MongoDB by their unique ObjectID.

Example:  
`DELETE /user/60c8d0f2e3b5c47a2c84f6d9`

#### 📤 Successful Response (200 OK)

```text
Deleted user with ID 60c8d0f2e3b5c47a2c84f6d9
```

---