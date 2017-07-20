package main

import (
	"./Router"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Starting server on :" + port)
	http.ListenAndServe(":"+port, Router.Routes())
}
