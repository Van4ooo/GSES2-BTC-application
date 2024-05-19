package tests

import (
	"GSES2_BTC_application/src/handlers"
	"GSES2_BTC_application/src/storage"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestGetRateHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/rate", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetRateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	responseBody := strings.TrimSpace(rr.Body.String())
	if _, err := strconv.Atoi(responseBody); err != nil {
		t.Errorf("handler returned non-numeric body: got %v, error: %v", responseBody, err)
	}
}

func TestSubscribeHandler(t *testing.T) {
	testCases := []struct {
		email           string
		expectedStatus  int
		expectedMessage string
	}{
		{"example1@gmail.com", http.StatusOK, "E-mail додано"},
		{"example2@gmail.com", http.StatusOK, "E-mail додано"},
		{"example1@gmail.com", http.StatusConflict, "Email already exists"},
		{"example2@gmail.com", http.StatusConflict, "Email already exists"},
		{"", http.StatusBadRequest, "Email is required"},
		{"sdgds.com", http.StatusBadRequest, "Email is required"},
	}

	for _, tc := range testCases {
		t.Run(tc.email, func(t *testing.T) {

			req, err := http.NewRequest("POST", "/api/subscribe", strings.NewReader("email="+tc.email))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.SubscribeHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
			}

			if strings.TrimSpace(rr.Body.String()) != tc.expectedMessage {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tc.expectedMessage)
			}
		})
	}
}

func TestSendEmailsHandler(t *testing.T) {
	t.Skip("Пропускаємо тестування /api/sendEmails," +
		" щоб воно не розсилало повідомлення користувачам")
	req, err := http.NewRequest("POST", "/api/sendEmails", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.SendEmailsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rr.Body.String() != "E-mailʼи відправлено" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "E-mailʼи відправлено")
	}
}

func TestEmailList(t *testing.T) {
	if _, err := storage.GetEmails(); err != nil {
		t.Errorf("GetEmails should return without error")
	}
}

func TestEmptyEmailsFile(t *testing.T) {
	if os.Remove(storage.EmailFile) != nil {
		t.Errorf("tests file [\"emails.txt\"] was not deleted")
	}

	if _, err := storage.GetEmails(); err == nil {
		t.Errorf("GetEmails should return an error")
	}
}
