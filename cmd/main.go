package main

import (
	"GSES2_BTC_application/src/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/rate", handlers.GetRateHandler).Methods("GET")
	router.HandleFunc("/api/subscribe", handlers.SubscribeHandler).Methods("POST")
	router.HandleFunc("/api/sendEmails", handlers.SendEmailsHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
