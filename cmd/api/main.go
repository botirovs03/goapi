package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/botirovs03/goapi/internal/handelers"]
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()

	handlers.Handler(r)

	fmt.Println("Starting server...")

	fmt.Println("GO API")

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Error(UnAuthorizedError)
		api.RequestErrorHandler(w, UnAuthorizedError)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.RequestErrorHandler(w)
		return
	}

	var loginDetails *tools.LoginDetails
	loginDetails = (*database).GetUserLoginDetails(username)

	if (loginDetails == nil || token != (*loginDetails).AuthToken) {
		log.Error(UnAuthorizedError)
		api.RequestErrorHandler(w, UnAuthorizedError)
		return
		
	}

	 next.ServeHTTP(w, r)
}