package helpers

import (
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 5)

func LimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// https://www.youtube.com/watch?v=qQBbCZTtboE&pp=ygUMZG94eGluZyBtZW1l
		log.Printf("%q: %q", r.RemoteAddr, r.UserAgent())

		// if timer isn't stopping yet - block request out
		if !limiter.Allow() {
			log.Printf("Too Many Requests By %q", r.RemoteAddr)
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		// if all goes well - go to handler
		next.ServeHTTP(w, r)
	})
}
