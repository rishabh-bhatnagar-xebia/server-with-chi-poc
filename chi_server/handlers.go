package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

var HandlePokemon http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pokemon"))
}

var LogRequest http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Path: %s\n", r.URL.Path)))
	w.Write([]byte(fmt.Sprintf("Vars: %s\n", chi.RouteContext(r.Context()).URLParams)))
}
