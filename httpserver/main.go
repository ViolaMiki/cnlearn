package main

import (
	"log"
	"net/http"
)

type Response struct {
	Status int
	http.ResponseWriter
}

func (r *Response) WriteHeader(code int) {
	r.Status = code
	r.ResponseWriter.WriteHeader(code)
}

func WithLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := &Response{
			Status:         http.StatusOK,
			ResponseWriter: w,
		}
		h.ServeHTTP(resp, r)
		log.Printf("Remote Addr %s , Http Status: %d", r.RemoteAddr,resp.Status)
	})
}

func main() {
	http.Handle("/header", WithLogging(http.HandlerFunc(headerHandler)))
	http.Handle("/version", WithLogging(http.HandlerFunc(versionHandler)))
	http.Handle("/healthz", WithLogging(http.HandlerFunc(healthzHandler)))
	http.ListenAndServe("127.0.0.1:8000", nil)
}