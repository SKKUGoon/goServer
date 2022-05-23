package broadcastWS

import (
	"fmt"
	"net/http"

	"log"
	"time"

	"github.com/gorilla/websocket"
	"goServer/broadcastStruct"
)

func (c *Client) readOrder() {
	defer func() {
		c.floor.unregister <- c
		c.conn.Close()
	}()

	var msgContain *broadcastStruct.MessageRecv
	for {
		// Read Sent JSON formatted string
		msgContain = &broadcastStruct.MessageRecv{}
		err := c.conn.ReadJSON(msgContain)
		if err != nil {
			log.Println(ClientJSONParseError, err)
			break
		}

		// Capture disconnection event
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf(ClientDisconnError, err)
			}
			break
		}

		c.floor.orderInfo <- msgContain
	}
}

func (c *Client) writeOrder() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.floor.unregister <- c
		c.conn.Close()
	}()

	for {
		select {
		// message handle
		case m, ok := <-c.call:
			fmt.Println("writeorder", m, ok)
			if !ok {
				// the hub closed the channel
				err := c.conn.WriteJSON(broadcastStruct.ConnCloseMessage)
				if err != nil {
					log.Println(ServerAbruptEnd)
				}
				return
			}

			//orderHandle(msg)

			err := c.conn.WriteJSON(broadcastStruct.TestMessage)
			log.Println(ServerMsgHand)
			if err != nil {
				log.Println(ServerSendingEnd)
				return
			}
		}
	}
}

func ServeWS(f *TradeFloor, w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Insert Upgraded HTTP connection(Websocket) into the TradingFloor
	client := &Client{
		floor: f,
		conn:  ws,
		call:  make(chan interface{}),
	}
	client.floor.register <- client

	go client.writeOrder()
	go client.readOrder()
}
