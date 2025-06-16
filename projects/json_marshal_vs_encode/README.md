# JSON Marshal vs Encode in Go

This is a simple Go web server that demonstrates how to convert Go structs to JSON using two methods:

- `json.Marshal` – Converts Go data into a JSON byte slice.
- `json.Encoder` – Converts Go data into JSON and writes it directly to an `io.Writer` (in this case, the HTTP response).

The server exposes three routes:

### 🔧 Routes

| Endpoint        | Description |
|----------------|-------------|
| `/`            | Serves a basic HTML page with a message. |
| `/marshal`     | Uses `json.Marshal` to convert a struct to JSON and sends it in the response. |
| `/encode`      | Uses `json.Encoder


### ▶️ How to Run
- Run the server  
```
go run main.go
```
- Open your browser or use curl to test the endpoints:
```
curl http://localhost:8080/marshal
curl http://localhost:8080/encode
```

### 🧠 Key Concepts
- `json.Marshal` gives you a JSON `[]byte` that you can manipulate or store.
- `json.NewEncoder(w).Encode()` writes JSON directly to the HTTP response, and is often cleaner for APIs.