package dateAndTime

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	apiUrl       = "http://localhost:8080"
	validRoute   = "/datetime"
	invalidRoute = "/datetimes"
)

func handleTestsErrors(t *testing.T, err error) {
	if err != nil {
		t.Errorf("An error occured %v", err)
		t.Fail()
	}
}

//To-DO : Server test 
//To-Do : mocking
//To-Do : benchmark testing

func TestDateAndTime(t *testing.T) {
	t.Run("Valid Api url", func(*testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			apiUrl+validRoute,
			nil,
		)
		resp := httptest.NewRecorder()
		GetDateAndTime(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", resp.Code, http.StatusOK)
		}

		var response Response
		err := json.NewDecoder(resp.Body).Decode(&response)
		handleTestsErrors(t, err)
		if response.DateTime != time.Now().Format("2006-01-02 15:04:05") {
			t.Errorf("expected date time format %s but got %s", time.Now().Format("2006-01-02 15:04:05"), response.DateTime)
		}
	})
	t.Run("Invalid Api url", func(*testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			apiUrl+invalidRoute,
			nil,
		)
		resp := httptest.NewRecorder()
		GetDateAndTime(resp, req)
		if resp.Code != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", resp.Code, http.StatusNotFound)
		}
	})
	t.Run("Invalid Method", func(*testing.T) {
		req := httptest.NewRequest(
			http.MethodPost,
			apiUrl+validRoute,
			nil,
		)
		resp := httptest.NewRecorder()
		GetDateAndTime(resp, req)
		if resp.Code != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", resp.Code, http.StatusMethodNotAllowed)
		}
	})
}
