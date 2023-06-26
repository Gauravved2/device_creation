package main

import (
	"net/http"

	"bitbucket.org/klokinnovations/webapp/service"
)

func main() {
	srv := service.NewServer()
	http.ListenAndServe(":8001",srv)
}