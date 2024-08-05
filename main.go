package main

import (
    "html/template"
    "net/http"
)

// Parse the template file
var templates = template.Must(template.ParseFiles("templates/index.html"))

// Handler function to render the HTML page
func handler(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "index.html", nil)
}

// Main function to set up the server
func main() {
    http.HandleFunc("/", handler)
    println("Server started at http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        println("Error starting server:", err)
    }
}
