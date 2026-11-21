package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/health", nil)
    recorder := httptest.NewRecorder()

    healthHandler(recorder, req)

    if recorder.Code != http.StatusOK {
        t.Errorf(
            "Expected status 200, got %d",
            recorder.Code,
        )
    }
}