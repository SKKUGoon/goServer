package main

import (
	"goServer/broadcast_ws"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Trading Session:")
	http.HandleFunc("/wss", broadcast_ws.Echo)
	log.Fatal(http.ListenAndServe(":7890", nil))
}
