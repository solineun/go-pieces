package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	terribleSecurity := TerribleSecurityProvoder("GOPHER")

	http.Handle("/hello", terribleSecurity(RequestTimer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello\n"))
		}))))
}

func RequestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("request time for %s: %v", r.URL.Path, end.Sub(start))
	})
}

var securityMsg = []byte("You didn't give the secret password\n")

func TerribleSecurityProvoder(password string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Secret-Password") != password {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(securityMsg)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}