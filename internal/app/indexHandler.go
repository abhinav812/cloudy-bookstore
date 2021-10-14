package app

import "net/http"

// HandleIndex - HandleIndex is the router to process index request.
func HandleIndex(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Length", "12")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte("Hello World!"))
}
