package network

import "net/http"

func Handler(
	inner func(w http.ResponseWriter, r *http.Request) *DurakResponse[any],
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		inner(w, r)
	}
}
