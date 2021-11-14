package middleware

import "net/http"

//ContextTypeJSON - router middleware to set content-type as JSON in every http response
func ContextTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json;charset=utf8")
		next.ServeHTTP(writer, request)
	})
}
