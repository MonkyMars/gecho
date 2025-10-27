package pkg

import (
	"net/http"
)

func OK(w http.ResponseWriter, payload *NewResponse) error {
	return WriteJSON(w, payload.Status, payload.Message, payload.Data)
}

func Err(w http.ResponseWriter, payload *NewResponse) error {
	return WriteJSON(w, payload.Status, payload.Message, nil)
}
