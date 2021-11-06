package main

import (
	"fmt"
	"net/http"

	"github.com/arturoeanton/go-analytics/pkg/handlers"
	"github.com/arturoeanton/go-analytics/pkg/u"
	"github.com/gorilla/mux"
)

func main() {
	u.LoadEnv(".env")
	port := u.EnvKey("PORT").Default("8080").String()

	r := mux.NewRouter()
	handlers.Register(r)
	fmt.Println("Starting server on port:", port)
	http.ListenAndServe(":"+port, r)
}
