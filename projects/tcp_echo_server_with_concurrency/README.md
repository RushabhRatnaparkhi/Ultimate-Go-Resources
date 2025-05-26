# TCP Echo Server with Concurrency

A concurrent TCP server in Go that listens on port `8080` and reads line-by-line input from connected clients. Each connection is handled in a separate goroutine, allowing the server to interact with multiple clients simultaneously.

---

## 📦 Features

- Listens for TCP connections on port `8080`
- Spawns a new goroutine for each client connection
- Reads client input line-by-line using `bufio.Scanner`
- Logs received input to the server console
- Demonstrates concurrent network programming in Go

---

## 🛠️ How It Works

- The server uses `net.Listen("tcp", ":8080")` to bind to port 8080
- Upon accepting a client connection, it launches a goroutine to handle input from that client
  ```
  go run main.go
  ```
- Input is read using a `bufio.Scanner`, which reads line-by-line
- Input is printed to the server console
- The connection is closed after the client disconnects
- You can simulate a client using the built-in nc (Netcat) tool:
  ```
  nc localhost 8080
  ```

---

## Output
![Screenshot of Output of TCP Echo Server with Concurrency](https://github.com/user-attachments/assets/3a788779-b901-4719-a3a1-3cad59defa22)
<img width="722" alt="Screenshot of localhost:8080 hit in browswer" src="https://github.com/user-attachments/assets/3472c806-705f-4a22-8985-e5bc5a30adc8" />

