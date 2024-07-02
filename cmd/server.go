package main

import (
	"fmt"
	dateAndTime "github.com/codescalersinternships/DateTime-Server-Mariam_Mahrous/pkg/http"
	"net/http"
)

func main() {
	
	http.HandleFunc(
		"/datetime",
		dateAndTime.GetDateAndTime,
	)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("an error occured")
	}
}
