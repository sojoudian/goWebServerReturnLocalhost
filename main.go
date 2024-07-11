package main

import (
    "fmt"
    "net/http"
    "os"
)

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err != nil {
        http.Error(w, "Unable to get hostname", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "Hostname: %s", hostname)
}

func main() {
    http.HandleFunc("/", hostnameHandler)
    fmt.Println("Server is running on port 80")
    if err := http.ListenAndServe(":80", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

