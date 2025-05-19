# 🧩 Template with Custom Functions to HTML

This example demonstrates how to use **custom functions** with Go's `text/template` package to enhance how data is displayed in an HTML template.

---

## 📦 What It Does

- Parses an HTML template file (`tpl.gohtml`)
- Defines two Go struct types:
  - `sage`: holds `Name` and `Motto`
  - `car`: holds `Manufacturer`, `Model`, and `Doors`
- Combines slices of these structs into one data structure
- Registers and uses custom template functions:
  - `uc`: Converts strings to UPPERCASE
  - `ft`: Returns the first 3 characters of a string
- Renders formatted HTML output to the terminal (or browser)

## 🚀 Run the Code

```
go run main.go
```

## ✅ When to Use

- When you want to transform or format data before displaying it in a template
- To add custom logic for text processing (e.g. trimming, masking, formatting)
- To keep your templates clean while centralizing logic in Go code

## References

- https://pkg.go.dev/text/template#Must
- https://pkg.go.dev/text/template#Template.Funcs

## Output

<img width="484" alt="Screenshot 2025-05-19 at 8 25 55 PM" src="https://github.com/user-attachments/assets/327531c6-7150-4cd4-9871-b0fd80e26b26" />
