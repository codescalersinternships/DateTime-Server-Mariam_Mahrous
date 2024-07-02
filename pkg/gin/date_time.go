package dateandtime

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type DateAndTime struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

func GetDateAndTime(c *gin.Context) {
	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now().Format("15:04:05")
	response := DateAndTime{Date: currentDate, Time: currentTime}
	c.IndentedJSON(http.StatusOK, response)
}
