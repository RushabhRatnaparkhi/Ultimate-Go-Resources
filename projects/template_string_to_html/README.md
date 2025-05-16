### Template string to html

This Go program demonstrates how to use the `text/template` package to parse and render template files to the terminal (standard output). It shows how to load multiple template files from a directory and render them individually.

```
go run main.go
```

## 🔍 Summary of Key Functions

| Function                             | Purpose                                                                 |
|--------------------------------------|-------------------------------------------------------------------------|
| `template.Must(...)`                 | Wraps template parsing and panics if there's an error.                 |
| `template.ParseGlob("templates/*")`  | Parses all matching template files in the specified folder.            |
| `tpl.Execute(os.Stdout, nil)`        | Executes the default (first parsed) template and prints to terminal.   |
| `tpl.ExecuteTemplate(w, name, data)` | Executes a specific named template and prints it to the writer.        |
| `os.Stdout`                          | Outputs the rendered template to the console.                          |


## 🛡️ Notes

- Use the `.gohtml` file extension for your templates.  
  This is a common Go convention and makes it easier to know which files are templates.

- If something goes wrong while loading a template (like a typo or missing file),  
  `template.Must` will stop the program right away and show an error.  
  This helps catch problems early before the program runs too far.
