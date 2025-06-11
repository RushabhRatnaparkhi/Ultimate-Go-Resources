package main

import (
	"crypto/sha1"   // Used to create a unique hash of the file contents
	"fmt"           // For formatted printing
	"html/template" // For parsing and executing HTML templates safely
	"io"            // For copying file content
	"net/http"      // For building the web server
	"os"            // For working with the file system
	"path/filepath" // For creating file paths that work across OSes
	"strings"       // For string operations

	uuid "github.com/satori/go.uuid" // Third-party library to generate unique session IDs
)

var tpl *template.Template // Global variable to hold parsed templates

// init runs before main and parses all templates in the 'templates/' folder
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)                          // Route for the home page
	http.Handle("/favicon.ico", http.NotFoundHandler()) // Avoid favicon request errors
	http.ListenAndServe(":8080", nil)                    // Start server on port 8080
}

// index handles GET and POST requests for "/"
func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req) // Get or create a user session cookie

	// Handle form submission
	if req.Method == http.MethodPost {
		// Retrieve uploaded file using the name of the file input field ("nf")
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close() // Ensure file stream is closed later

		// Extract file extension (e.g., "jpg", "png") from original filename
		ext := strings.Split(fh.Filename, ".")[1]

		// Create a SHA1 hash of the file content to generate a unique filename
		h := sha1.New()
		io.Copy(h, mf) // Read file content into the hash
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext // Hash + original extension

		// Determine the path where the file will be saved
		wd, err := os.Getwd() // Get current working directory
		if err != nil {
			fmt.Println(err)
		}
		// Combine path: projectDir/public/pics/hashedFilename.ext
		path := filepath.Join(wd, "public", "pics", fname)

		// Create the new file on disk
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		// Reset file reader to the beginning and copy content into the new file
		mf.Seek(0, 0)
		io.Copy(nf, mf)

		// Update the session cookie to include this new file name
		c = appendValue(w, c, fname)
	}

	// Split the cookie value to get all uploaded filenames
	xs := strings.Split(c.Value, "|")

	// Pass the list of filenames to the HTML template for rendering
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

// getCookie retrieves the session cookie or creates a new one if missing
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		// If cookie doesn't exist, generate a new UUID
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",      // Cookie name
			Value: sID.String(),   // Unique session ID
		}
		http.SetCookie(w, c) // Set the new cookie in the user's browser
	}
	return c
}

// appendValue appends a new filename to the cookie value if it's not already present
func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname // Append using "|" as a separator
	}
	c.Value = s
	http.SetCookie(w, c) // Update the cookie in the browser
	return c
}
