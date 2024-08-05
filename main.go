package main

import (
    "html/template"
    "net/http"
)

var templates = template.Must(template.ParseFiles("templates/index.html"))

func handler(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "index.html", nil)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        name := r.FormValue("name")
        greeting := "Hi, " + name + "! Nice to meet you!"
        w.Write([]byte(greeting))
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/greet", greetHandler)
    println("Server started at http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        println("Error starting server:", err)
    }
}
