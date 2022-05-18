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
| pongWait = 60sec         |
| pingPeriod = 54sec       |
+--------------------------+
| Serving on = localhost   |
+--------------------------+
| Dependency: Gorilla      |
| Some middle agents       |
| written in python        |
+--------------------------+
| LICENCE:                 |
+--------------------------+
`

	ServerNewConnection = "[Trading Floor Assistance] >>> new client connected"
	ServerEndConnection = "[Trading Floor Assistance] >>> client unreachable. disconnecting"
)

// Server Error Messages for log
const (
	ServerAbruptEnd  = "[Trading Floor Errors] >>> closing message sending error"
	ServerSendingEnd = "[Trading Floor Errors] >>> sending response error"
)

const (
	ServerMsgRecv = "[Trading Floor Comm] >>> incoming new message"
	ServerMsgHand = "[Trading Floor Comm] >>> new message handled"
)
