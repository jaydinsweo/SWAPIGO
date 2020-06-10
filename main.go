package main

import (
	"log"
	"net/http"

	apiroute "github.com/jaynguyens/SWAPIGO/app"
)

func main() {
	router := apiroute.AppRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
