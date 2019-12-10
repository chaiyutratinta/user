package main

import (
	"log"
	"net/http"
	"user/router"
)

func main() {
	route := router.New()
	log.Fatal(http.ListenAndServe(":8080", route))
}
