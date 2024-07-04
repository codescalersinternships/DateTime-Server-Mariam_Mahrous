package dateandtime

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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

func TestDateAndTime(t *testing.T) {
	r := gin.Default()
	r.GET("/datetime", GetDateAndTime)
	t.Run("Valid Api url", func(*testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			apiUrl+validRoute,
			nil,
		)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", resp.Code, http.StatusOK)
		}

		var response DateAndTime
		err := json.NewDecoder(resp.Body).Decode(&response)
		handleTestsErrors(t, err)
		maxAcceptedTime := time.Now().Add(0 + 0 + time.Second*time.Duration(3)).Format("15:04:05")

		if response.Date != time.Now().Format("2006-01-02") && response.Time >= time.Now().Format("15:04:05") && response.Time <= maxAcceptedTime {
			t.Errorf("expected date %s but got %s , time %s but got %s ", time.Now().Format("2006-01-02"), response.Date, time.Now().Format("15:04:05"), response.Time)
		}
	})
	t.Run("invalid Api url", func(*testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			apiUrl+invalidRoute,
			nil,
		)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		if resp.Code != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", resp.Code, http.StatusNotFound)
		}
	})
	t.Run("invalid Method", func(*testing.T) {
		req := httptest.NewRequest(
			http.MethodPost,
			apiUrl+validRoute,
			nil,
		)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		if resp.Code != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", resp.Code, http.StatusNotFound)
		}
	})
}
