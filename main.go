package main

import "net/http"

const message = "Microservice B1210C is online !"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})

	http.ListenAndServe(":8081", mux)
}
