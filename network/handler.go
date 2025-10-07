package network

import "net/http"

func Handler(
	inner func(w http.ResponseWriter, r *http.Request) *DurakHandlerResult,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		inner(w, r)
	}
}
