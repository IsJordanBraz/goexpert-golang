package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	http.ListenAndServe(":9000", mux)
}
