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
}

func signalHandle(s MessageRecv, cl Client) {
	switch {
	case s.SignalType == "init":
		clientInit(s, cl.Connect)

	case s.SignalType == "conn":
		fmt.Println("Signal conn")
		cl.connResp()

	case s.SignalType == "spot_trade":
		fmt.Println("Signal spot trade")
		cl.tradeResp(false)

	case s.SignalType == "spread_trade":
		fmt.Println("Signal spread trade")
		cl.tradeResp(false)

	case s.SignalType == "test_trade":
		fmt.Println("Signal test_trade")
		cl.tradeResp(true)
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
		signalHandle(*m, c)
	}
}

func (c Client) connResp() {
	// Responding to conn heartbeat
}

func (c Client) tradeResp(isTest bool) {
	// Responding to trade messages
	if isTest == true {
		fmt.Println("Test Signal")
	} else {
		fmt.Println("Real Trade Signal")
	}
}
