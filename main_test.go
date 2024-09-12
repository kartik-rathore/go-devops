package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := `<h1 id="greeting">Hi, my name is Kartik Rathore</h1>`
    if !strings.Contains(rr.Body.String(), expected) {
        t.Errorf("handler returned unexpected body: got %v want to contain %v", rr.Body.String(), expected)
    }
}

func TestGreetHandler(t *testing.T) {
    tests := []struct {
        name           string
        method         string
        formData       string
        expectedStatus int
        expectedBody   string
    }{
        {"Valid POST", "POST", "name=John", http.StatusOK, "Hi, John! Nice to meet you!"},
        {"Invalid Method", "GET", "", http.StatusMethodNotAllowed, "Invalid request method\n"},
        {"Empty Name", "POST", "name=", http.StatusBadRequest, "Name is required\n"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req, err := http.NewRequest(tt.method, "/greet", strings.NewReader(tt.formData))
            if err != nil {
                t.Fatal(err)
            }
            if tt.method == "POST" {
                req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
            }

            rr := httptest.NewRecorder()
            handler := http.HandlerFunc(greetHandler)

            handler.ServeHTTP(rr, req)

            if status := rr.Code; status != tt.expectedStatus {
                t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
            }

            if rr.Body.String() != tt.expectedBody {
                t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expectedBody)
            }
        })
    }
}
