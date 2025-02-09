package main

import (
	"log"
	"net/http"
	"time"
)

// LogRequest logs incoming HTTP requests
func LogRequest(r *http.Request) {
	log.Printf("[%s] %s %s from %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, r.RemoteAddr)
}
