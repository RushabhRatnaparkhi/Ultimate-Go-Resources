# 🧩 Template Safe HTML Rendering

This project demonstrates how to safely render dynamic HTML content using Go's `html/template` package.  
The output is written directly to an `index.html` file.

## ⚙️ How It Works

- Go reads a template file `template.gohtml`.
- A `Page` struct holds user data, including potentially unsafe input.
- The template engine safely escapes this input to prevent XSS attacks.
- The rendered HTML is written to `index.html`.

## 🚀 How to Run

Make sure you're in the `template_safe_html_rendering/` directory and run:

```bash
go run main.go > index.html
```

## Output

![Template Safe HTML Rendering Screenshot](https://github.com/user-attachments/assets/98271de5-c6b6-419d-a176-51ad65510119)
