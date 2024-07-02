package dateAndTime

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DateAndTime struct {
	Date string `json:"date"`
	Time string `json:"time"`
}


func GetDateAndTime(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/datetime" {
		http.Error(w, fmt.Sprintf("invalid Request URI: %s", r.RequestURI), http.StatusNotFound)
		return
	} else if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("invalid http method: %s", r.Method), http.StatusMethodNotAllowed)
		return
	}

	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now().Format("15:04:05")
	response := DateAndTime{Date: currentDate , Time: currentTime}
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
