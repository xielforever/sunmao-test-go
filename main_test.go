package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	mux := newMux("1.2.3")
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var body healthResp
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("decode failed: %v", err)
	}
	if body.Status != "ok" || body.Service != "sunmao-test-go" || body.Version != "1.2.3" {
		t.Fatalf("unexpected health body: %+v", body)
	}
}

func TestInfo(t *testing.T) {
	mux := newMux("1.2.3")
	req := httptest.NewRequest(http.MethodGet, "/api/info", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var body infoResp
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("decode failed: %v", err)
	}
	if body.Service != "sunmao-test-go" || body.Version != "1.2.3" {
		t.Fatalf("unexpected info body: %+v", body)
	}
}
