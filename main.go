package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(map[string]string{
        "status":  "ok",
        "service": "helpdesk-api",
    })
}

func main() {
    http.HandleFunc("GET /health", healthHandler)

    log.Println("Server running on port 8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}