package main

import (
	"fmt"
	"net/http"

	"github.com/dansonserge/golang-with-danson/go-bscs/cmd/apiDev/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)
func main() {
	log.SetReportCaller(true)
	//create a multiplexer of chi

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting go API Server...")

	err := http.ListenAndServe("localhost:8000",r)

	if err != nil {
		log.Error(err)
	}

}