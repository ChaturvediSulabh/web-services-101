package cors

import (
	"net/http"
)

//Middleware ...
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Allow-Access-Control-Methods", "POST,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length")
		h.ServeHTTP(w, r)
	})
}
