package broadcast_ws

import (
	"github.com/gorilla/websocket"
	"log"
)

// TODO: create struct

func JsonRecv(conn *websocket.Conn) {
	m := &MessageRecv{}

	err := conn.ReadJSON(m)
	if err != nil {
		log.Println("disconnetion?", err)
		return
	} else {
		// Client is On board
		cl := Client{
			Connect:  conn,
			MyOffice: m.Data.Dep,
		}
		signalHandle(*m, cl)
	}
}

func JsonResp(conn *websocket.Conn) {
	r := MessageResp{
		SignalType: "conn_resp",
		Data: DataResp{
			Status: "normal",
			Msg:    "connection_normal",
		},
	}
	err := conn.WriteJSON(r)
	if err != nil {
		log.Println(err)
		return
	}
}

func sigHandle(s MessageRecv, conn *websocket.Conn) interface{} {
	switch {
	case s.SignalType == "init":
		addClient(s, conn)

	case s.SignalType == "conn":
		return nil

	case s.SignalType == "spot_trade":
		return nil
	case s.SignalType == "spread_trade":
		return nil
	case s.SignalType == "test_trade":
		return nil
	}
	return nil
}

func addClient(m MessageRecv, conn *websocket.Conn) {
	switch {
	case m.Data.Dep == "back":
		BackOfficeGroup[conn] = true
		log.Println(CastBack0)

	case m.Data.Dep == "midl":
		MiddOfficeGroup[conn] = true
		log.Println(CastMidd0)

	case m.Data.Dep == "frnt":
		FrntOfficeGroup[conn] = true
		log.Println(CastFrnt0)

	default:
		log.Println(CastBack1)
	}
}
