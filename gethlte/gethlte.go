package gethlte

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/ayoseun/geth-lte/internals"
	"github.com/ayoseun/geth-lte/internals/address"
	"github.com/ayoseun/geth-lte/internals/block"
	"github.com/ayoseun/geth-lte/internals/mempool"
	"github.com/ayoseun/geth-lte/internals/transactions"
	"github.com/ayoseun/geth-lte/types"
	"github.com/gorilla/websocket"
)

func GetBalance(rpc string, walletAddress string) ([]byte, error) {
	//"https://bsc.meowrpc.com"
	walletBalance, err := address.GetAddressBalance(rpc, walletAddress)

	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err
	}

	response := types.WalletBalanceResponse{
		Balance: walletBalance,
		Address: walletAddress,
	}

	// Marshal the response to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err
	}

	// Return the JSON-encoded response as a []byte
	return responseJSON, nil

}

func GetAddressTransactionCount(rpc string, walletAddress string) (string, error) {
	//"https://bsc.meowrpc.com"
	txCount, err := address.GetAddressTXCount(rpc, walletAddress)

	if err != nil {
		// Return an error if there's a problem fetching the balance

	}

	return txCount, nil

}

func GetGasPrice(rpc string) (string, error) {
	//"https://bsc.meowrpc.com"
	walletBalance, err := internals.GasPrice(rpc)

	if err != nil {
		// Return an error if there's a problem fetching the balance

	}

	// Return the JSON-encoded response as a []byte
	return walletBalance, nil

}

func GetLatestBlock(rpc string) (string, error) {
	//"https://bsc.meowrpc.com"
	block, err := block.LastBlock(rpc)

	if err != nil {
		// Return an error if there's a problem fetching the balance

	}

	// Return the JSON-encoded response as a []byte
	return block, nil

}

func GetBlockTransactionCount(rpc string, hash string) (string, error) {
	//"https://bsc.meowrpc.com"
	txCount, err := block.GetBlockTXCountByHash(rpc, hash)

	if err != nil {
		// Return an error if there's a problem fetching the balance

	}

	return txCount, nil

}

// GetTransactionByHash retrieves a transaction by its hash using the provided RPC URL.
func GetTransactionByHash(rpc string, hash string) ([]byte, error) {
	// Call the address.GetTransactionByHash function to fetch the transaction data
	tx, err := transactions.GetTransactionByHash(rpc, hash)
	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err

	}

	txData := types.TransactionResponse{
		JSONRPC: tx.JSONRPC,
		ID:      tx.ID,
		Result:  tx.Result,
	}
	// Marshal the response to JSON
	responseJSON, err := json.Marshal(txData)
	if err != nil {

		// Return an error if there's a problem fetching the balance
		return nil, err
	}

	// Return the JSON-encoded response as a []byte

	return responseJSON, nil
}

// GetTransactionByHash retrieves a transaction by its hash using the provided RPC URL.
func GetTransactionReceiptByHash(rpc string, hash string) ([]byte, error) {
	// Call the address.GetTransactionByHash function to fetch the transaction data
	tx, err := transactions.GetTransactionReceiptByHash(rpc, hash)
	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err

	}

	txData := types.TransactionRecieptResponse{
		JSONRPC: tx.JSONRPC,
		ID:      tx.ID,
		Result:  tx.Result,
	}
	// Marshal the response to JSON
	responseJSON, err := json.Marshal(txData)
	if err != nil {

		// Return an error if there's a problem fetching the balance
		return nil, err
	}

	// Return the JSON-encoded response as a []byte

	return responseJSON, nil
}

// GetTransactionByHash retrieves a transaction by its hash using the provided RPC URL.
func GetBlockByHash(rpc string, hash string) ([]byte, error) {
	// Call the address.GetTransactionByHash function to fetch the transaction data
	block, err := block.GetBlockByHash(rpc, hash)
	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err

	}

	blockData := types.BlockResponse{
		JSONRPC: block.JSONRPC,
		ID:      block.ID,
		Result:  block.Result,
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

// TransactionsMemPool retrieves transaction data from the mempool using the provided RPC URL.
func TransactionsMemPool(rpc string) ([]byte, error) {
	// Call the mempool.TxMemPool function to fetch the transaction data
	pool, err := mempool.TxMemPool(rpc)
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
	pool, err := mempool.TxMemPoolCount(rpc)
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

		ch <- string(message)
	}

	// Handle Ctrl+C to gracefully close the WebSocket connection
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	fmt.Println("Closing WebSocket connection...")

}