package main

import (
	"fmt"
	ginserver "github.com/codescalersinternships/DateTime-Server-Mariam_Mahrous/pkg/gin"
	httpserver "github.com/codescalersinternships/DateTime-Server-Mariam_Mahrous/pkg/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//http
	http.HandleFunc(
		"/datetime",
		httpserver.GetDateAndTime,
	)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("an error occured")
	}

	//ginserver
	router := gin.Default()
	router.GET("./datetime", ginserver.GetDateAndTime)
	router.Run("localhost:8000")
}
