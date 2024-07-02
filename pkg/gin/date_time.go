package dateandtime

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// A struct that has Time and Date
type DateAndTime struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

// An api that returns the current date and time
func GetDateAndTime(c *gin.Context) {
	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now().Format("15:04:05")
	response := DateAndTime{Date: currentDate, Time: currentTime}
	c.IndentedJSON(http.StatusOK, response)
}
