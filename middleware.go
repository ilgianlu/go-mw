package middleware

import (
	"fmt"
	"net/http"
)

//LogMiddleware ... ciao bello
func LogMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	   fmt.Println("OK....")
	   h.ServeHTTP(w, r) // call ServeHTTP on the original handler
	})
 }
