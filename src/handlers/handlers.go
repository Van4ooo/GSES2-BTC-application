package handlers

import (
	"GSES2_BTC_application/src/email"
	"GSES2_BTC_application/src/storage"
	"GSES2_BTC_application/src/third-party-api"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
)

func GetRateHandler(w http.ResponseWriter, r *http.Request) {
	rate, err := third_party_api.GetRateBTC_UAH()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rate)
}

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println()
	_email := r.FormValue("email")
	fmt.Println(_email)
	_, err := mail.ParseAddress(_email)

	if _email == "" || err != nil {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	if storage.IsEmailExists(_email) {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	if err := storage.AddEmail(_email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("E-mail додано"))
}

func SendEmailsHandler(w http.ResponseWriter, r *http.Request) {
	rate, err := third_party_api.GetRateBTC_UAH()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	emails, err := storage.GetEmails()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = email.SendEmails(emails, rate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("E-mailʼи відправлено"))
}
