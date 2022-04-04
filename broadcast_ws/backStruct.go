package broadcast_ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Connect  *websocket.Conn
	MyOffice string
}

type ClientAction interface {
	MsgRecv()
	connResp()
	tradeResp()
	testResp()
}

func signalHandle(s MessageRecv, conn *websocket.Conn) {
	switch {
	case s.SignalType == "init":
		clientInit(s, conn)
	case s.SignalType == "conn":
		fmt.Println("Signal conn")
	case s.SignalType == "spot_trade":
		fmt.Println("Signal spot trade")
	case s.SignalType == "spread_trade":
		fmt.Println("Signal spread trade")
	case s.SignalType == "test_trade":
		fmt.Println("Signal test_trade")
	}
}

func clientInit(mr MessageRecv, conn *websocket.Conn) {
	switch {
	case mr.Data.Dep == "back":
		BackOfficeGroup[conn] = true
		log.Println(CastBack0)
	case mr.Data.Dep == "midl":
		MiddOfficeGroup[conn] = true
		log.Println(CastMidd0)
	case mr.Data.Dep == "frnt":
		FrntOfficeGroup[conn] = true
		log.Println(CastFrnt0)
	}
}

func (c Client) MsgRecv() {
	m := &MessageRecv{}

	err := c.Connect.ReadJSON(m)
	if err != nil {
		log.Println(err)
	} else {
		signalHandle(*m, c.Connect)
	}
}

func (c Client) connResp() {
	// Responding to conn heartbeat
}

func (c Client) tradeResp() {
	// Responding to trade messages
}

func (c Client) testResp() {
	// Responding to test trade messages
}
