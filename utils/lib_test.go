package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestExtractResponseBody(t *testing.T) {
	w := httptest.NewRecorder()
	if err := NewOK(w, WithData(map[string]string{"extract": "test"})).Send(); err != nil {
		t.Fatal(err)
	}
	value, err := ExtractResponseBody[NewResponse](w.Result())
	if err != nil || value.Data().(map[string]any)["extract"] != "test" {
		t.Fatalf("value=%v err=%v", value, err)
	}
}

func TestWriteJSON(t *testing.T) {
	w := httptest.NewRecorder()
	if err := writeJSON(w, http.StatusTeapot, true, "tea", nil, map[string]string{"tea": "yes"}); err != nil {
		t.Fatal(err)
	}
	var value NewResponse
	if err := json.NewDecoder(w.Body).Decode(&value); err != nil || value.Status() != http.StatusTeapot {
		t.Fatalf("value=%v err=%v", value, err)
	}
}

func TestWriteJSONNilWriter(t *testing.T) {
	if err := writeJSON(nil, http.StatusOK, true, "x", nil, nil); err == nil {
		t.Fatal("expected nil writer error")
	}
}

func TestGetTimestamp(t *testing.T) {
	before := time.Now()
	got := getTimestamp()
	if got.Before(before) || got.After(time.Now().Add(time.Second)) {
		t.Fatalf("timestamp out of range: %v", got)
	}
}

func TestNewResponseBuilder(t *testing.T) {
	w := httptest.NewRecorder()
	resp := NewOK(w)
	if resp == nil || w.Code != 200 {
		t.Fatalf("resp=%v code=%d", resp, w.Code)
	}
}
