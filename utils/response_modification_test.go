package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponseModification(t *testing.T) {
	w := httptest.NewRecorder()
	resp := NewOK(w, WithData(map[string]string{"initial": "value"}))
	resp.SetMessage("Modified").SetStatus(http.StatusCreated).AddData("added", true)
	if err := resp.Send(); err != nil {
		t.Fatal(err)
	}
	var body NewResponse
	if err := json.NewDecoder(w.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	data := body.Data().(map[string]any)
	if body.Message() != "Modified" || data["initial"] != "value" || data["added"] != true {
		t.Fatalf("unexpected response: %#v", body)
	}
}

func TestAddDataConversions(t *testing.T) {
	tests := []struct {
		name string
		data any
	}{
		{"nil", nil}, {"string map", map[string]string{"old": "value"}},
		{"int map", map[string]int{"old": 1}}, {"other", "value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := NewOK(httptest.NewRecorder(), WithData(tt.data))
			if resp.AddData("new", true) == nil {
				t.Fatal("AddData returned nil")
			}
			if _, ok := resp.data.(map[string]any); !ok {
				t.Fatalf("data type = %T", resp.data)
			}
		})
	}
}

func TestNilResponseMethods(t *testing.T) {
	var resp *Response
	if resp.SetMessage("x") != nil || resp.SetStatus(200) != nil ||
		resp.SetData("x") != nil || resp.AddData("x", true) != nil {
		t.Fatal("nil response mutation should remain nil")
	}
	if err := resp.Send(); err != nil {
		t.Fatal(err)
	}
}
