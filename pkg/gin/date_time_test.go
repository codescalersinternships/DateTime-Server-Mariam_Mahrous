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
	apiUrl     = "http://localhost:8080"
	validRoute = "/datetime"
	//invalidRoute = "/datetimes"
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
	r := gin.Default()
	t.Run("Valid Api url", func(*testing.T) {
		r.GET("/datetime", GetDateAndTime)
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

		if response.Date != time.Now().Format("2006-01-02") && response.Time != time.Now().Format("15:04:05") {
			t.Errorf("expected date %s but got %s , time %s but got %s ", time.Now().Format("2006-01-02"), response.Date, time.Now().Format("15:04:05"), response.Time)
		}
	})
}
