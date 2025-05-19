# Template Struct Data to HTML

This example demonstrates how to pass structured data (Go structs and slices) from a Go program into an HTML template using Go's `text/template` package.

## What It Does

- Parses the template file `template.gohtml`
- Defines two Go struct types:
  - `animal`: contains fields like `Name` and `Sound`
  - `habitat`: contains fields like `Location`, `Type`, and `Climate`
- Passes slices of these structs into the template
- Renders dynamic lists in the HTML using Go template syntax
- Outputs the result to the terminal (or a browser if used with an HTTP server)

## 🚀 Run the Code

```bash
go run main.go
```

## ✅ When to Use

This pattern is useful when:

- Displaying structured or nested data (like lists, records, etc.)
- Creating dynamic reports, dashboards or summaries in HTML
- Generating templates from Go structs (for documentation, output files, etc.)

## Output
<img width="713" alt="Template Struct Data to HTML Screenshot" src="https://github.com/user-attachments/assets/43a9e77d-7841-49ef-9497-719ed7a46d6f" />

