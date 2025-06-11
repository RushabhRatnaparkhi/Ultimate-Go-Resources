package main

import (
	"crypto/sha1"   // For creating a unique hash of the uploaded file
	"fmt"           // For formatted output and logging
	"html/template" // For rendering dynamic HTML templates
	"io"            // For reading and writing file streams
	"net/http"      // For building the web server and handling HTTP requests
	"os"            // For accessing the file system (reading current directory, creating files)
	"path/filepath" // For building file paths in a platform-independent way
	"strings"       // For string manipulation (splitting and checking cookie values)

	uuid "github.com/satori/go.uuid" // External package to generate UUIDs for session IDs
)

var tpl *template.Template // Global variable to hold parsed HTML templates

// init function is automatically called before main().
// It parses all templates in the 'templates/' directory and stores them in 'tpl'.
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// Register the index handler for the root URL
	http.HandleFunc("/", index)

	// Serve static files (like uploaded images) from the "public" directory
	// Accessed via URLs like /public/pics/filename.jpg
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	// Ignore requests to favicon.ico (prevents unnecessary errors)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}

// index handles both GET (render page) and POST (handle file upload) requests
func index(w http.ResponseWriter, req *http.Request) {
	// Get or create a session cookie
	c := getCookie(w, req)

	// Handle file upload if it's a POST request
	if req.Method == http.MethodPost {
		// Retrieve the uploaded file and its metadata from the form
		mf, fh, err := req.FormFile("nf") // "nf" is the name of the form field
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close() // Ensure file stream is closed when done

		// Extract file extension from the original filename (e.g., "jpg", "png")
		ext := strings.Split(fh.Filename, ".")[1]

		// Create a SHA1 hash of the file contents to use as a unique filename
		h := sha1.New()
		io.Copy(h, mf) // Read the file into the hash function
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext // Format as hex string with original extension

		// Get the current working directory to construct the full path
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		// Create the full path for saving the uploaded file
		path := filepath.Join(wd, "public", "pics", fname)

		// Create a new file at the target path
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		// Reset the read pointer of the uploaded file to the beginning
		// (we already read it for hashing above)
		mf.Seek(0, 0)

		// Copy the contents of the uploaded file to the new file on disk
		io.Copy(nf, mf)

		// Append the new filename to the user's session cookie value
		c = appendValue(w, c, fname)
	}

	// Split the cookie value by "|" to get list of uploaded filenames
	xs := strings.Split(c.Value, "|")

	// xs[0] is the UUID (session ID), so we skip it and pass only file names to the template
	tpl.ExecuteTemplate(w, "index.gohtml", xs[1:])
}

// getCookie tries to read the session cookie.
// If it doesn't exist, it creates a new UUID-based session cookie.
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		// Cookie doesn't exist, so generate a new session ID
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",     // Name of the cookie
			Value: sID.String(),  // Unique session ID
		}
		// Send the new cookie to the user's browser
		http.SetCookie(w, c)
	}
	return c
}

// appendValue adds a new filename to the cookie if it's not already present
func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	// Only append if the filename isn't already in the cookie (avoid duplicates)
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c) // Update the cookie in the browser
	return c
}
