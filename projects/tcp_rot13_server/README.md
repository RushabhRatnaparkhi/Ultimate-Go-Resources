# TCP ROT13 Transformation Server

A concurrent TCP server written in Go that listens on port `8080`, accepts input from clients, converts it to lowercase, applies a ROT13 transformation, and sends both the original and transformed strings back to the client.

---

## 📦 Features

- Listens for TCP connections on port `8080`
- Handles multiple client connections concurrently using goroutines
- Converts all input to lowercase before processing
- Applies the ROT13 cipher transformation (basic letter rotation by 13)
- Sends both the original and encoded response back to the client

---

## 🛠️ How It Works

- The server listens on TCP port `8080` using `net.Listen`.
- For each new client, a separate goroutine is launched to handle communication.
- It reads lines from the client using `bufio.Scanner`.
- Each line is:
  - Converted to lowercase
  - Transformed using a ROT13 cipher
  - Responded with the format:  
    ```
    <original> - <rot13 transformed>
    ```

---

## Output

<img width="894" alt="Screenshot of Output of TCP rot13 transformation server" src="https://github.com/user-attachments/assets/548cd597-a2fe-4105-9f20-8bb7118e64de" />

