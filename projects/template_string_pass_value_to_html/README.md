# Template String Pass Value to HTML

This example demonstrates how to pass a string value from a Go program into an HTML template using Go's `text/template` package.

## What It Does

- Parses all `.gohtml` files inside the `templates` directory
- Executes two templates:
  - **`template.gohtml`**: Injects a single string value (`"awesome"`)
  - **`looptemplate.gohtml`**: Injects a slice of strings (`["good", "bad", "ugly"]`)
- Uses Go template syntax to insert the data into the HTML
- Renders the output to the terminal (or a browser if connected to a web server)


## 🚀 Run the Code

```bash
go run main.go
```

## ✅ When to Use

This pattern is helpful when:

- Rendering dynamic content into HTML pages
- Creating email templates
- Generating text or configuration files dynamically

## Output

![Screenshot of Template String Pass Value to HTMLM](https://github.com/user-attachments/assets/4d1ffdc5-7816-4fc2-82b9-a6c181ef63e9)

--------

 **Template Package** : https://pkg.go.dev/text/template#pkg-index
