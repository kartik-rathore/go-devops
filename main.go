package main

import (
    "html/template"
    "log"
    "net/http"
)

var templates = template.Must(template.ParseFiles("templates/index.html"))

func handler(w http.ResponseWriter, r *http.Request) {
    if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    name := r.FormValue("name")
    if name == "" {
        http.Error(w, "Name is required", http.StatusBadRequest)
        return
    }

    greeting := "Hi, " + name + "! Nice to meet you!"
    if _, err := w.Write([]byte(greeting)); err != nil {
        log.Printf("Error writing response: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/greet", greetHandler)
    log.Println("Server started at http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
