package broadcastWS

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type TradeFloor struct {
	// connected workers
	traders map[*Client]bool

	// broadcast chan
	orderInfo chan interface{}

	// register clients
	register chan *Client

	// unregister clients
	unregister chan *Client
}

type Client struct {
	// Where Clients are in
	floor *TradeFloor

	// The connection itself
	conn *websocket.Conn

	// Buffered channel message
	call chan interface{}
}

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
			fmt.Println(ServerNewConnection)
			fl.traders[client] = true

		case client := <-fl.unregister:
			if _, ok := fl.traders[client]; ok {
				// if the client is in trader close it
				fmt.Println(ServerEndConnection)

				// close client's message sending channel
				delete(fl.traders, client)
				close(client.call)
			}

		case message := <-fl.orderInfo:
			fmt.Println(ServerMsgRecv, message)
			for client := range fl.traders {
				select {
				// push message into call channel of the client
				case client.call <- message:

				// if fail: client is malfunctioning. delete client
				default:
					close(client.call)
					delete(fl.traders, client)
				}
			}
		}
	}
}
