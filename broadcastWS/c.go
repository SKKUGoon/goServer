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

// Server Connection messages
const (
	NewTradingSession = `
##### [Trading Floor] #####
+------- Parameters -------+
| pongWait   = 60sec       |
| pingPeriod = 54sec       |
| *ActivePingPong disabled |
+------- Service on -------+
| Serving on = localhost   |
| PortNum    = :7890       |
+------- Dependency -------+
| Dependency: Gorilla      |
| Some middle agents       |
| written in python        |
+------- Additional -------+
| LICENCE: MIT LICENSE     |
+--------------------------+
`
	ServerNewConnection = "[Trading Floor Assistance] >>> new client connected"
	ServerEndConnection = "[Trading Floor Assistance] >>> client unreachable. disconnecting"
)

// Server Error Messages for log
const (
	TradingFloorMsgError = "[Trading Floor Errors] >>> client malfunctioned. deleting client"
	ServerAbruptEnd      = "[Trading Floor Errors] >>> closing message sending error"
	ServerSendingEnd     = "[Trading Floor Errors] >>> sending response error"
)

const (
	ServerMsgRecv = "[Trading Floor Comm] >>> incoming new message"
	ServerMsgHand = "[Trading Floor Comm] >>> new message handled"
)

// Client-side messages
const (
	ClientJSONParseError = "[Client Errors] >>> JSON message parsing error"
	ClientDisconnError   = "[Client Errors] >>> %v"
)
