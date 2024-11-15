package main

import (
    "log"
    "net/http"
    "retail-pulse-backend/handlers"
    "retail-pulse-backend/store"
)

func main() {
    // Load store master data
    if err := store.LoadStoreMasterData("store_master.json"); err != nil {
        log.Fatalf("Failed to load store master data: %v", err)
    }

    // Setup routes
    http.HandleFunc("/api/submit", handlers.SubmitJob)
    http.HandleFunc("/api/status", handlers.GetJobStatus)

    // Start server
    log.Println("Starting server on port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
