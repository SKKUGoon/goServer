package broadcastWS

import "github.com/gorilla/websocket"

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
