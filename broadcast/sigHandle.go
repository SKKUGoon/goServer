package broadcast

import (
	"github.com/gorilla/websocket"
	"log"
)

func JsonRecv(conn *websocket.Conn) {
	m := &MessageRecv{}

	err := conn.ReadJSON(m)
	if err != nil {
		log.Println(err)
	} else {
		sigHandle(*m, conn)
	}
}

func JsonResp(conn *websocket.Conn) {

}

func sigHandle(s MessageRecv, conn *websocket.Conn) interface{} {
	switch {
	case s.SignalType == "init":
		addClient(s, conn)
	case s.SignalType == "trade":
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
