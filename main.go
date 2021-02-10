package main

import (
	"fmt"
	"net/http"
)

func mpHandler(urlMap map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := urlMap[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		} else if r.URL.Path == "/" {
			fmt.Fprintf(w, "<h1>Welcome to the homepage</h1>")
		}
	}
}

func main() {
	mux := http.NewServeMux()

	urlPaths := map[string]string{"/goog": "https://google.com", "/fb": "https://facebook.com"}

	mph := mpHandler(urlPaths)
	mux.HandleFunc("/", mph)

	http.ListenAndServe(":3000", mux)
}
