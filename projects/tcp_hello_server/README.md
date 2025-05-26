# TCP Hello Server

A simple TCP server written in Go that listens on port `8080` and sends a friendly message to any client that connects. This is a minimal example demonstrating how to use Go's `net` package to accept TCP connections, write responses, and handle basic I/O.

## 🛠️ How It Works

Upon running the server:
- It opens a TCP listener on `localhost:8080`
```
go run main.go
```
- Accepts incoming client connections (e.g., via `telnet` or `nc`)
```
nc localhost 8080
```

## Output
<img width="1288" alt="Output Screenshot of tcp hello server" src="https://github.com/user-attachments/assets/ae1a738d-4aef-401e-96d5-6146b9c8a5ce" />
