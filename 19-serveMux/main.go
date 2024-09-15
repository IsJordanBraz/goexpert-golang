package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home"))
	})
	mux.HandleFunc("/hello", handleHello)
	mux.Handle("/blog", blog{title: "my blog"})

	http.ListenAndServe(":8080", mux)
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
