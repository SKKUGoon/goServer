package broadcastStruct

var MsgConnResp = MessageResp{
	SignalType: "conn_resp",
	Data: DataResp{
		Status: "normal",
		Msg:    "connection_normal",
	},
}

var TestMessage = MessageResp{
	SignalType: "conn_resp",
	Data: DataResp{
		Status: "normal",
		Msg:    "connection normal from websocket",
	},
}

var ConnCloseMessage = MessageResp{
	SignalType: "conn_resp",
	Data: DataResp{
		Status: "closing",
		Msg:    "connection closed, bye",
	},
}
