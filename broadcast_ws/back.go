package broadcast_ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"goServer/broadcast_struct"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ReadWrite(conn *websocket.Conn) {
	fmt.Println("3")
	var msgContain *broadcast_struct.MessageRecv
	for {
		msgContain = &broadcast_struct.MessageRecv{}

		// Read Sent JSON formatted string
		err := conn.ReadJSON(msgContain)
		if err != nil {
			log.Println("disconnection event", err)
			break
		}

		fmt.Println("recv", msgContain)
		fmt.Println("send", broadcast_struct.TestMessage)

		// Write JSON formatted Response

		/*
			실험하기
			1. 두 개의 Seperate Connection 연결시 일어나는 일.
			2. 백오피스나 Client Struct 안에 저장할때 일어나는 일.
		*/
		err = conn.WriteJSON(broadcast_struct.TestMessage)
		if err != nil {
			log.Println("sending error", err)
			break
		}

	}
}

func Echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Serve Func", err)
		return
	}
	fmt.Println("2")
	defer conn.Close()
	fmt.Println("1")
	ReadWrite(conn)
}
