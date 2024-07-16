package main

import (
	"fmt"
	ginserver "github.com/codescalersinternships/DateTime-Server-Mariam_Mahrous/pkg/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("./datetime", ginserver.GetDateAndTime)
	err := router.Run(":8000")
	if err != nil {
		fmt.Printf("an error occured")
	}
}
