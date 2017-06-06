package main

import (
	"net/http"

	"github.com/GolangAce/experiment/common"
	"github.com/GolangAce/experiment/handlers"
	"github.com/gorilla/mux"
)

func main() {
	mainSession := common.NewMongoSession()
	defer mainSession.Close()

	r := mux.NewRouter()
	http.Handle("/", r)
	r.Handle("/", handlers.IndexHandler(mainSession))
	http.ListenAndServe(":8080", r)
}
