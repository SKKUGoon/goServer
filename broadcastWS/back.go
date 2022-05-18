package broadcastWS

import (
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

	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { return c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	var msgContain *broadcastStruct.MessageRecv
	for {
		msgContain = &broadcastStruct.MessageRecv{}

		// Read Sent JSON formatted string
		err := c.conn.ReadJSON(msgContain)

		// Capture disconnection event
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
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
		case _, ok := <-c.call:
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

		// ping pong handle
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServeWSS(f *TradeFloor, w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		floor: f,
		conn:  ws,
		call:  make(chan interface{}),
	}
	client.floor.register <- client

	go client.writeOrder()
	go client.readOrder()
}
