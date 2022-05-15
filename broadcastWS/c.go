package broadcastWS

import (
	"github.com/gorilla/websocket"
	"time"
)

const (
	writeWait = 10 * time.Second

	// Define Websocket ping-pong interval
	pongWait   = 60 * time.Second
	pingPeriod = pongWait * 9 / 10
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
