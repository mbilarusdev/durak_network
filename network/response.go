package network

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type DurakResponse[T any] struct {
	Code    int        `json:"code"`
	Data    T          `json:"data,omitempty"`
	Err     *ErrorPart `json:"err,omitempty"`
	Jwt     string
	Headers map[string]string
}

type DurakHandlerResult struct{}

func (r *DurakResponse[T]) Send(w http.ResponseWriter) *DurakHandlerResult {
	responseJson, err := json.Marshal(r)
	if err != nil {
		log.Println("Ошибка маршалинга HTTP-ответа")
	}
	w.Header().Set("Content-Type", "application/json")
	if r.Jwt != "" {
		w.Header().Set("Authorization", fmt.Sprintf("Bearer %v", r.Jwt))
	}
	for name, value := range r.Headers {
		w.Header().Set(name, value)
	}
	w.WriteHeader(int(r.Code))
	w.Write(responseJson)
	return &DurakHandlerResult{}
}
