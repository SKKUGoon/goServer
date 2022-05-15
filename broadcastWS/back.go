package broadcastWS

import (
	"fmt"
	"net/http"

	"log"
	"time"

	"github.com/gorilla/websocket"
	"goServer/broadcastStruct"
)

func NewFloor() *TradeFloor {
	return &TradeFloor{
		traders:    make(map[*Client]bool),
		orderInfo:  make(chan interface{}),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (fl *TradeFloor) RunFloor() {
	for {
		select {
		case client := <-fl.register:
			fl.traders[client] = true
		case client := <-fl.unregister:
			if _, ok := fl.traders[client]; ok {
				// if the client is in trader close it
				// close client's message sending channel
				delete(fl.traders, client)
				close(client.call)
			}
		case message := <-fl.orderInfo:
			for client := range fl.traders {
				select {
				case client.call <- message:
				default:
					close(client.call)
					delete(fl.traders, client)
				}
			}
		}
	}
}

func (c *Client) readOrder() {
	defer func() {
		// to unregister
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
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.call:
			fmt.Println("msgmsg", msg)
			fmt.Println("ok", ok)
			if !ok {
				// the hub closed the channel
				fmt.Println("the hub closed the channel")
				return
			}
			fmt.Println("ws", msg)

			err := c.conn.WriteJSON(broadcastStruct.TestMessage)
			if err != nil {
				log.Println("sending error, wss", err)
				return
			}
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
