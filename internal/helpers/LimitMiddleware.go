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
		log.Println(r.RemoteAddr)

		if !limiter.Allow() {
			log.Println("Too Many Requests By " + r.RemoteAddr)
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
