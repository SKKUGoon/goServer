package main

import (
	"flag"
	"goServer/broadcastWS"
	"log"
	"net/http"
)

//func main() {
//	log.Println("Starting Trading Session:")
//	http.HandleFunc("/wss", broadcastWS.Echo)
//	log.Fatal(http.ListenAndServe(":7890", nil))
//}

var addr = flag.String("addr", ":7890", "http service address")

func main() {
	fl := broadcastWS.NewFloor()
	go fl.RunFloor()

	http.HandleFunc("/wss", func(w http.ResponseWriter, r *http.Request) {
		broadcastWS.ServeWSS(fl, w, r)
	})
	err := http.ListenAndServe(":7890", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
