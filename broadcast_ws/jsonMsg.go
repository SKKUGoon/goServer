package broadcast_ws

var MsgConnResp = MessageResp{
	SignalType: "conn_resp",
	Data: DataResp{
		Status: "normal",
		Msg:    "connection_normal",
	},
}
