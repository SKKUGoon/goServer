package main

import (
	"fmt"
	"goServer/broadcastWS"
	"log"
	"net/http"
)

func init() {
	fmt.Printf(broadcastWS.NewTradingSession)
}

func main() {
	fl := broadcastWS.NewFloor()
	go fl.RunFloor()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		broadcastWS.ServeWS(fl, w, r)
	})

	err := http.ListenAndServe(":7890", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
