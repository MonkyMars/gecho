package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkResponseSend(b *testing.B) {
	for b.Loop() {
		resp := NewOK(httptest.NewRecorder(), WithData(map[string]any{"id": 1, "name": "user"}))
		if err := resp.Send(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkImmediateResponseSend(b *testing.B) {
	for b.Loop() {
		if err := NewOK(httptest.NewRecorder(), WithData(map[string]any{"id": 1, "name": "user"})).Send(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNoContentSend(b *testing.B) {
	for b.Loop() {
		if err := NewOK(httptest.NewRecorder(), WithStatus(http.StatusNoContent)).Send(); err != nil {
			b.Fatal(err)
		}
	}
}
