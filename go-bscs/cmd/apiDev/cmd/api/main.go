package main

import (
	"github.com/dansonserge/apidev/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)
func main() {

	log.SetReportCaller(true)

	//create a multiplexer of chi

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

}