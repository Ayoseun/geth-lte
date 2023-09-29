package mempool

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"github.com/ayoseun/geth-lte/internals/mempool/mempool_core"
	"github.com/ayoseun/geth-lte/types"
	"github.com/gorilla/websocket"
)

// TransactionsMemPool retrieves transaction data from the mempool using the provided RPC URL.
func TransactionsMemPool(rpc string) ([]byte, error) {
	// Call the mempool.TxMemPool function to fetch the transaction data
	pool, err := mempool_core.TxMemPool(rpc)
	if err != nil {
		// Return an error if there's a problem fetching the transaction data
		return nil, err
	}
	// Create a new MemPoolData struct and populate it with data from the 'pool' variable
	poolData := types.MemPool{
		JSONRPC: pool.JSONRPC,
		ID:      pool.ID,
		Result:  pool.Result,
	}
	// Marshal the transaction data to JSON
	responseJSON, err := json.Marshal(poolData)
	if err != nil {
		// Return an error if there's a problem marshaling to JSON
		return nil, err
	}

	// Return the JSON-encoded transaction data as []byte
	return responseJSON, nil
}

// GetTransactionByHash retrieves a transaction by its hash using the provided RPC URL.
func TransactionsMemPoolCount(rpc string) ([]byte, error) {
	// Call the address.GetTransactionByHash function to fetch the transaction data
	pool, err := mempool_core.TxMemPoolCount(rpc)
	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err

	}

	blockData := types.MemPoolCount{
		Result: pool.Result,
	}
	// Marshal the response to JSON
	responseJSON, err := json.Marshal(blockData)
	if err != nil {

		// Return an error if there's a problem fetching the balance
		return nil, err
	}

	// Return the JSON-encoded response as a []byte

	return responseJSON, nil
}

func RTLTransactionsMempool(wssURL string, ch chan string) {
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

	// Create an EthSubscriptionData object
	var ethData EthSubscriptionData

	// Unmarshal the message into the EthSubscriptionData struct
	if err := json.Unmarshal(message, &ethData); err != nil {
		log.Println("Error unmarshaling message:", err)
		continue
	}

	// Marshal the EthSubscriptionData struct back to JSON
	ethDataJSON, err := json.Marshal(ethData.Params.Result)
	if err != nil {
		log.Println("Error marshaling EthSubscriptionData:", err)
		continue
	}

	// Send the JSON data to the channel as []byte
	ch <- string(ethDataJSON)
		// You can now access the data using ethSubscription
		// For example, you can access the blockHash as ethSubscription.params.result.blockHash

		// Send the JSON data to the channel
		//ch <- string(message)
	}

	// Handle Ctrl+C to gracefully close the WebSocket connection
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	fmt.Println("Closing WebSocket connection...")
}

type EthSubscriptionResult struct {
	BlockHash        *string `json:"blockHash"`
	BlockNumber      *string `json:"blockNumber"`
	From             string  `json:"from"`
	Gas              string  `json:"gas"`
	GasPrice         string  `json:"gasPrice"`
	Hash             string  `json:"hash"`
	Input            string  `json:"input"`
	Nonce            string  `json:"nonce"`
	To               string  `json:"to"`
	TransactionIndex *string `json:"transactionIndex"`
	Value            string  `json:"value"`
	Type             string  `json:"type"`
	V                string  `json:"v"`
	R                string  `json:"r"`
	S                string  `json:"s"`
}

type EthSubscriptionParams struct {
	Result EthSubscriptionResult `json:"result"`
}

type EthSubscriptionData struct {
	JSONRPC     string              `json:"jsonrpc"`
	Method      string              `json:"method"`
	Params      EthSubscriptionParams `json:"params"`
	Subscription string              `json:"subscription"`
}
