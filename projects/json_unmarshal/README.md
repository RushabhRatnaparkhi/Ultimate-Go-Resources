# JSON Unmarshal in Go

This simple Go program demonstrates how to **parse JSON data into Go structs** using `json.Unmarshal`. It also explains the importance of using a **pointer** (`&data`) to let the `Unmarshal` function modify the original variable.

### 🧠 Key Concepts
- `json.Unmarshal(data, &v)` takes a JSON byte slice and a pointer to a Go variable.
- The pointer (`&v`) tells Go where to store the parsed data.
- Without the pointer, `Unmarshal` would only modify a copy of the struct and your data wouldn't be saved.

### ▶️ How to Run
```
go run main.go
```