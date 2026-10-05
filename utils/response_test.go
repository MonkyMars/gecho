package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewOKBuildsAndSends(t *testing.T) {
	w := httptest.NewRecorder()
	resp := NewOK(w, WithMessage("Test Message"), WithData(map[string]string{"key": "value"}), WithStatus(http.StatusAccepted))
	if err := resp.Send(); err != nil {
		t.Fatal(err)
	}
	if w.Code != http.StatusAccepted {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusAccepted)
	}
	var body NewResponse
	if err := json.NewDecoder(w.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body.Message() != "Test Message" || !body.Success() {
		t.Fatalf("unexpected response: message=%q success=%v", body.Message(), body.Success())
	}
}

func TestNewErr(t *testing.T) {
	w := httptest.NewRecorder()
	resp := NewErr(w, WithMessage("Error Message"), WithData("Error Details"), WithStatus(http.StatusBadRequest))
	if err := resp.Send(); err != nil {
		t.Fatal(err)
	}
	var body NewResponse
	if err := json.NewDecoder(w.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body.Success() || body.Data() != "Error Details" {
		t.Fatalf("unexpected error response: success=%v data=%v", body.Success(), body.Data())
	}
}

func TestResponseOptionsAndHeaders(t *testing.T) {
	w := httptest.NewRecorder()
	resp := NewOK(w, WithHeader("X-Test", "one"), WithHeaders(map[string]string{"X-Test": "two", "X-Other": "value"}))
	if err := resp.Send(); err != nil {
		t.Fatal(err)
	}
	if w.Header().Get("X-Test") != "two" || w.Header().Get("X-Other") != "value" {
		t.Fatalf("headers = %v", w.Header())
	}
}

func TestWithoutSendDoesNotWrite(t *testing.T) {
	w := httptest.NewRecorder()
	_ = NewOK(w, WithData("not sent"))
	if w.Body.Len() != 0 {
		t.Fatalf("body length = %d, want 0", w.Body.Len())
	}
}

func TestNilWriterError(t *testing.T) {
	if err := NewOK(nil).Send(); err == nil {
		t.Fatal("expected nil writer error")
	}
}
