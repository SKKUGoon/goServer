package broadcast

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Connect  *websocket.Conn
	MyOffice string
}

type ClientAction interface {
	MsgRecv()
	MsgResp()
}

func (c Client) MsgRecv() {

}

func (c Client) MsgResp() {

}
