package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	mux := http.NewServeMux()
	// // server := http.DefaultServeMux
	// r := mux.Handle()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, r.URL.Path)
		t := time.Now()

		time.Since(t)

	})
	// r := mux.NewRouter().StrictSlash(false)
	// http.Handle("",http.StripPrefix(".",http.NotFoundHandler()))
	// r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	// r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	// http.ListenAndServe(":8080", mux)
}
