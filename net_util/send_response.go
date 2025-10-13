package net_util

import (
	"encoding/json"
	"log"
	"net/http"

	net_res "github.com/mbilarusdev/durak_network/response"
)

func SendResponse(
	w http.ResponseWriter,
	statusCode int,
	body *net_res.Body,
	headers map[string]string,
) {
	jsonBody, err := json.Marshal(body)

	if err != nil {
		log.Println("Ошибка маршалинга ответа")
		return
	}

	for key, value := range headers {
		w.Header().Set(key, value)
	}

	w.Write(jsonBody)
	w.WriteHeader(statusCode)
}
