package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

// TestHandler checks the response from the main handler function
func TestHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handler)

    handler.ServeHTTP(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body
    expected := `<h1 id="greeting">Hi, my name is Kartik Rathore</h1>`
    if !strings.Contains(rr.Body.String(), expected) {
        t.Errorf("handler returned unexpected body: got %v want to contain %v", rr.Body.String(), expected)
    }
}

// TestGreetHandler checks the response from the greet handler function
func TestGreetHandler(t *testing.T) {
    // Test valid POST request
    form := "name=John"
    req, err := http.NewRequest("POST", "/greet", strings.NewReader(form))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(greetHandler)

    handler.ServeHTTP(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body
    expected := "Hi, John! Nice to meet you!"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }

    // Test invalid request method
    req, err = http.NewRequest("GET", "/greet", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr = httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusMethodNotAllowed {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
    }

    // Check the response body
    expected = "Invalid request method\n"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
