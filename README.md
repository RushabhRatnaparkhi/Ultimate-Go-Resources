# Ultimate Go Resources

<div align="center">
  
[![Join Our Discord](https://img.shields.io/badge/Discord-Join%20Server-blue?logo=discord&style=for-the-badge)](https://discord.com/invite/Yn9g6KuWyA)
[![Subscribe on YouTube](https://img.shields.io/badge/YouTube-Subscribe-red?logo=youtube&style=for-the-badge)](https://www.youtube.com/@dhanushnehru?sub_confirmation=1)
[![Subscribe to Newsletter](https://img.shields.io/badge/Newsletter-Subscribe-orange?style=for-the-badge)](https://dhanushn.substack.com/)

</div>

<img width="1600" height="840" alt="Ultimate Golang Resources" src="https://github.com/user-attachments/assets/2364b0ce-f189-4311-8441-d5f0aa648eb1" />

<a href="https://www.producthunt.com/products/ultimate-golang-resources?embed=true&utm_source=badge-featured&utm_medium=badge&utm_source=badge-ultimate&#0045;golang&#0045;resources" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=992783&theme=light&t=1752676381336" alt="Ultimate&#0032;Golang&#0032;Resources - A&#0032;curated&#0032;list&#0032;of&#0032;collection&#0032;of&#0032;Go&#0032;projects&#0032;and&#0032;resources&#0032; | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>

# 📖 Documentation
- 🔗 Official GoLang Docs: [Go Dev](https://go.dev/)
- [Go Tour](https://go.dev/tour/welcome/1)
- [Go Packages](https://pkg.go.dev/)

## 📁 Available Projects
- [cards](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/cards): A simple Go program that defines a new custom type deck (a slice of strings) representing a deck of playing cards. It demonstrates creating, populating, and printing the deck using nested loops and receiver functions.
- [string_to_html](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/string_to_html): Converts a plain string into an HTML document using Go and writes the output to a file.
- [template_string_to_html](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/template_string_to_html): Defines an HTML template as a raw string in the Go file and renders it with inserted data.
- [template_pass_string_value_to_html](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/template_pass_string_value_to_html): Passes a string value into an external HTML template and renders it dynamically using Go templates.
- [template_struct_data_to_html](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/template_struct_data_to_html): Demonstrates how to pass structured data (using Go structs and slices) into an external HTML template and render dynamic content using Go’s text/template package.
- [template_with_custom_functions_to_html](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/template_with_custom_functions_to_html): Demonstrates how to register and use custom functions in Go templates to manipulate and display structured data dynamically.
- [template_with_pipelining](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/template_with_pipelining): Demonstrates using function pipelines in Go templates to transform and render data step-by-step with custom functions.
- [template_safe_html_rendering](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/template_safe_html_rendering): Demonstrates safe rendering of dynamic user input in Go templates using the `html/template` package, which automatically escapes potentially dangerous content to prevent XSS attacks.


- [tcp_hello_server](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/tcp_hello_server): A simple TCP server in Go that listens on port 8080 and sends a greeting message to any client that connects. Demonstrates basic use of the `net` package for creating TCP servers and handling connections.
- [tcp_echo_server_with_concurrency](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/tcp_echo_server_with_concurrency): A concurrent TCP server in Go that listens on port 8080 and reads text input from clients line-by-line. Demonstrates handling multiple client connections using goroutines and reading data from network streams with `bufio.Scanner`.
- [tcp_echo_response_server](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/tcp_echo_response_server): A concurrent TCP server in Go that listens on port 8080, reads input from clients line-by-line, and sends back a custom echoed response. Demonstrates bidirectional communication using `bufio.Scanner` and `fmt.Fprintf`, along with concurrent connection handling using goroutines.

- [tcp_rot13_server](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/tcp_rot13_server): A concurrent TCP server in Go that listens on port 8080, reads input from clients, converts it to lowercase, and sends back a ROT13-transformed response. Demonstrates text transformation, bidirectional communication using `bufio.Scanner` and `fmt.Fprintf`, and concurrent client handling using goroutines.

- [photo_blog](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/photo_blog): A simple web application in Go that allows users to upload image files via an HTML form, which are saved with a SHA1-hashed filename. Uploaded images are stored in the `public/pics` directory and displayed back to the user based on session-tracked cookies.

- [hmac_hash](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/hmac_hash): A simple Go program that demonstrates generating an HMAC using SHA-256. It takes input strings (e.g., email addresses) and returns a secure, consistent hash using a secret key. Illustrates the use of `crypto/hmac`, `crypto/sha256`, and `io.WriteString` for keyed hashing.
- [base64_encoder](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/base64_encoder): A simple Go program that demonstrates Base64 encoding and decoding of a plain text message using Go's `encoding/base64` package. Shows how to convert strings to Base64 format and decode them back, with basic error handling using `log.Fatalln`.

- [json_marshal_vs_encode](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/json_marshal_vs_encode): A simple Go web server that demonstrates the difference between `json.Marshal` and `json.Encoder`. It defines a struct and shows how to convert it to JSON using both methods. Includes endpoints to serve basic HTML, marshal data to JSON, and encode data directly into the HTTP response, with proper headers and error handling.
- [json_unmarshal](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/json_unmarshal): A simple Go program that demonstrates how to parse a JSON string into nested Go structs using `json.Unmarshal`. It shows how to access nested fields, iterate over slices, and explains the importance of passing a pointer (`&data`) so the unmarshal operation modifies the original variable. Ideal for beginners learning how to handle JSON in Go.
- [crud_mongodb](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/crud_mongodb): A basic CRUD API built with Go and MongoDB that demonstrates how to connect to a MongoDB instance using the official Go driver. The project follows clean structuring with models, controllers, and routing logic, making it a great starting point for developers looking to build RESTful services with MongoDB as the database layer.
- [crud_postgres](https://github.com/DhanushNehru/Ultimate-Go-Resources/tree/main/projects/crud_postgres): A simple CRUD API built with Go and PostgreSQL that demonstrates how to connect to a PostgreSQL database using the `database/sql` package and `lib/pq` driver. The project uses a clean folder structure with separate `models` and `controllers`, and includes standard RESTful routes for creating, reading, updating, and deleting user records. Ideal for developers who want to learn how to build robust backend services with PostgreSQL as the relational data layer.


### Connect With Me

Let's stay connected! Follow me on social media to stay updated:

[![X](https://img.shields.io/badge/Twitter-Follow-blue?style=flat-square&logo=X)](https://x.com/Dhanush_Nehru) 
[![Instagram](https://img.shields.io/badge/Instagram-Follow-blue?style=flat-square&logo=instagram)](https://www.instagram.com/dhanush_nehru/) 
[![YouTube](https://img.shields.io/badge/YouTube-Subscribe-red?style=flat-square&logo=youtube)](https://www.youtube.com/@dhanushnehru?sub_confirmation=1) 
[![GitHub](https://img.shields.io/badge/GitHub-Follow-blue?style=flat-square&logo=github)](https://github.com/DhanushNehru)


### Gitpod

In the cloud-free development environment where you can directly start coding.

You can use Gitpod in the cloud  [![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/DhanushNehru/Ultimate-Go-Resources/)

----

Feel free to update the README.md or raise issues if any to enhance the project
