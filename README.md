# WebServer for All Asset

<p>
    
This application uses `gorilla/websocket` 
and delivers message from client to client
    
</p>

### Scheme

<p>

1. Trading Floor is generated
2. HTTP HandleFunc with callbacks are defined.
3. Serve WebServer
4. If anytime a new connection is tried, `ServeWS` function will be called.

ServeWS is played out like so

```go
func ServeWS(f *TradeFloor, w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Insert Upgraded HTTP connection(Websocket) into the TradingFloor
	client := &Client{
		floor: f,
		conn:  ws,
		call:  make(chan interface{}),
	}
	client.floor.register <- client

	go client.writeOrder()
	go client.readOrder()
}
```

1. HTTP TCP Handshake is activated.
2. TCP Handshake is upgraded into websocket connection via parameters defined in `upgrader`
3. New Client structure <em>"C"</em>  is generated with necessary go channel
4. New Client <em>"C"</em> is registered into TradingFloor
5. Start <em>"C"</em> `writeOrder` and `readOrder` go routine

</p>

### TODO
<p>

* Create message wise handle function

</p>