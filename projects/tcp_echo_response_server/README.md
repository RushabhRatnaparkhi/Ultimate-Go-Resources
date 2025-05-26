# TCP Echo Response Server

A TCP server in Go that listens on port `8080`, reads client input line-by-line, and sends an echoed response back. Each connection is handled in its own goroutine, allowing the server to support multiple clients concurrently.

---

## 📦 Features

- Listens for incoming TCP connections on port `8080`
- Handles multiple clients concurrently using goroutines
- Reads text input line-by-line using `bufio.Scanner`
- Echoes back each line to the client with a custom message
- Logs all received input to the server console

---

## 🛠️ How It Works

- The server starts by listening on TCP port `8080` using `net.Listen`.
- When a client connects, a new **goroutine** is spawned to handle that client.
- The server reads data from the client **line-by-line** using `bufio.Scanner`.
- For each line:
  - The server prints it to the terminal
  - Sends a response like: `I heard you say: <client input>`
- The connection is closed automatically when the client disconnects or ends the stream.

---

## Output

<img width="1070" alt="Screenshot Output of tcp echo response server" src="https://github.com/user-attachments/assets/4461d00c-f8a4-43fb-8041-bfbcb863afd5" />
