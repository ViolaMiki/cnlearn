package main

import (
	"net/http"
	"os"
)

func headerHandler(rw http.ResponseWriter, req *http.Request) {
	for key, header := range req.Header {
		rw.Header().Add(key, header[0])
	}
}

func versionHandler(rw http.ResponseWriter, req *http.Request) {
	envs:= os.Getenv("VERSION")

	rw.Header().Add("VERSION", envs)
}

func healthzHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("ok"))
}