package main

import (
	"log"
	"net/http"
	"time"
)

func SleeperMiddleware(sleepFor time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Middleware starting to sleep on:", time.Now())
			time.Sleep(sleepFor)
			log.Println("Middleware woke up by", time.Now())

			next.ServeHTTP(w, r)
		})
	}
}
