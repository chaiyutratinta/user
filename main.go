package main

import (
	"log"
	"net/http"
	"user/configs"
	"user/router"
)

func main() {
	route := router.New()

	log.Fatal(http.ListenAndServe(configs.Config.Server.Port, route))
}
