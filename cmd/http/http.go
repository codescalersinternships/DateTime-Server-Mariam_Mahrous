package main

import (
	"fmt"
	httpserver "github.com/codescalersinternships/DateTime-Server-Mariam_Mahrous/pkg/http"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/datetime",
		httpserver.GetDateAndTime,
	)
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("server is running on port 8080")
	if err != nil {
		fmt.Printf("an error occured")
	}
}
