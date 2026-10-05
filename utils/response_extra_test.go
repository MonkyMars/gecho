package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewResponseJSONRoundTrip(t *testing.T) {
	when := time.Date(2026, time.January, 2, 3, 4, 5, 0, time.UTC)
	original := NewResponse{status: 201, success: true, message: "created", timestamp: when}
	encoded, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), `"data"`) {
		t.Fatalf("nil data should be omitted: %s", encoded)
	}
	var decoded NewResponse
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	if !decoded.Timestamp().Equal(when) || decoded.Status() != 201 {
		t.Fatalf("round trip mismatch")
	}
}

func TestNoContentHasEmptyBody(t *testing.T) {
	w := httptest.NewRecorder()
	if err := NewOK(w, WithStatus(http.StatusNoContent)).Send(); err != nil {
		t.Fatal(err)
	}
	if w.Code != http.StatusNoContent || w.Body.Len() != 0 {
		t.Fatalf("status=%d body=%q", w.Code, w.Body.String())
	}
}

func TestExtractResponseBodyErrors(t *testing.T) {
	if _, err := ExtractResponseBody[NewResponse](&http.Response{Body: http.NoBody}); err == nil {
		t.Fatal("expected empty body error")
	}
}

func TestResponseBuilderDoesNotWriteDuringConstruction(t *testing.T) {
	w := httptest.NewRecorder()
	resp := NewOK(w)
	if resp == nil || w.Body.Len() != 0 {
		t.Fatal("constructor wrote a response")
	}
}
