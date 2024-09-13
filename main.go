package main

import (
	"log"
	"net/http"
	handler "ritikjainrj18/tail/Handler"
)

func main() {

	http.HandleFunc("/log", handler.LogWatchHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/test", handler.TestHandler)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
