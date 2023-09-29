package mempool_core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/ayoseun/geth-lte/types"
	"github.com/gorilla/websocket"
)

func WebsocketMempool(wssURL string, ch chan string) {

	// Connect to the WebSocket server
	u, err := url.Parse(wssURL)
	if err != nil {
		log.Fatal(err)
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	request := types.JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_subscribe",
		Params:  []interface{}{"alchemy_newFullPendingTransactions"},
		ID:      1,
	}

	// Marshal the request to JSON
	requestJSON, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	// Send the JSON-RPC request
	if err := c.WriteMessage(websocket.TextMessage, requestJSON); err != nil {
		log.Fatal(err)
	}

	// Receive and handle JSON-RPC response
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}

		ch <- string(message)
	}

	// Handle Ctrl+C to gracefully close the WebSocket connection
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	fmt.Println("Closing WebSocket connection...")

}
