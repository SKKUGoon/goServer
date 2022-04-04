package main

import (
	"goServer/broadcast_ws"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Trading Session:")
	broadcast_ws.SetupRoutes()
	log.Fatal(http.ListenAndServe(":7890", nil))
}
