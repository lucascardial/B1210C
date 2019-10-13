package main

import (
	"B1210C/server"
	"log"
	"net/http"
	"os"
)

var (
	GcukCertFile    = os.Getenv("GCUK_CERT_FILE")
	GcukKeyFile     = os.Getenv("GCUK_KEY_FILE")
	GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
)

const message = "Microservice B1210C is online !"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	srv := server.New(mux, GcukServiceAddr)

	err := srv.ListenAndServeTLS(GcukCertFile, GcukKeyFile)

	if err != nil {
		log.Fatalf("SERVER FAILED TO START: %v", err)
	}
}
