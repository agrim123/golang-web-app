package main

import (
	"./Database"
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
	Database.Connect()
	// Database.ParseConfigFile()
	fmt.Println("Starting server on :" + port)
	http.ListenAndServe(":"+port, Router.Routes())
}
