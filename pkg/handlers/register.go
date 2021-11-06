package handlers

import "github.com/gorilla/mux"

func Register(r *mux.Router) {
	r.HandleFunc("/", HomeHandler)
}
