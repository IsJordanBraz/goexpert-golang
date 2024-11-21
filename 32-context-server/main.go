package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer log.Println("Request ended")
	select {
	case <-time.After(time.Second * 3):
		log.Println("Done")
		w.Write([]byte("Hello"))
	case <-r.Context().Done():
		log.Println("Cancel")
	}
}
