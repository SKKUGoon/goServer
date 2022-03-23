package main

import (
	"goServer/broadcast"
	"log"
	"net/http"
	//"net/http"
)

func main() {
	log.Println("Starting Trading Session:")
	broadcast.SetupRoutes()
	log.Fatal(http.ListenAndServe(":7890", nil))
}
