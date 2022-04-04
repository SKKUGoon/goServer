package broadcast_ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	BackOfficeGroup = make(map[*websocket.Conn]bool)
	MiddOfficeGroup = make(map[*websocket.Conn]bool)
	FrntOfficeGroup = make(map[*websocket.Conn]bool)
)

func homePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Home Page")
	if err != nil {
		return
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")
	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		// Define MessageRecv Structure
		JsonRecv(conn)

		// Send MessageRecv
		for c := range BackOfficeGroup {
			JsonResp(c)
		}
	}
}

func SetupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}
