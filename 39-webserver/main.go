package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /books/{id}", GetBookHandler)

	// books/dir/123/daw/123/vawd
	mux.HandleFunc("GET /books/dir/{d...}", BooksPathHandler)

	// books - EXATO
	mux.HandleFunc("GET /books/{$}", BooksHandler)

	// books/precedence/latest
	mux.HandleFunc("GET /books/precedence/latest", BooksPrecedenceHandler)

	// books/precedence/aaaaa
	mux.HandleFunc("GET /books/precedence/{x}", BooksPrecedence2Handler)

	// books/latest - conflicts
	mux.HandleFunc("GET /books/{x}", BooksPrecedenceHandler)

	// books/latest - conflicts
	mux.HandleFunc("GET /{x}/latest", BooksPrecedence2Handler)

	http.ListenAndServe(":9000", mux)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	dirpath := r.PathValue("d")
	fmt.Fprintf(w, "Accessing directory path: %s\n", dirpath)
}

func BooksPathHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Book: " + r.PathValue("id")))
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books"))
}

func BooksPrecedenceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence"))
}

func BooksPrecedence2Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence 2"))
}
