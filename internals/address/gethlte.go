package address

import (
	"encoding/json"

	"github.com/ayoseun/geth-lte/internals"
	"github.com/ayoseun/geth-lte/internals/address/address_core"
	"github.com/ayoseun/geth-lte/types"
	"github.com/ayoseun/geth-lte/types/results"
)

func GetBalance(rpc string, walletAddress string) ([]byte, error) {
	//"https://bsc.meowrpc.com"
	walletBalance, err := address_core.GetAddressBalance(rpc, walletAddress)

	if err != nil {
		// Return an error if there's a problem fetching the balance
		return nil, err
	}

	response := results.WalletBalanceResponse{
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
	txCount, err := address_core.GetAddressTXCount(rpc, walletAddress)

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
// GetTransactionByHash retrieves a transaction by its hash using the provided RPC URL.
func GetTransactionByHash(rpc string, hash string) ([]byte, error) {
	// Call the address.GetTransactionByHash function to fetch the transaction data
	tx, err := address_core.GetTransactionByHash(rpc, hash)
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
	tx, err := address_core.GetTransactionReceiptByHash(rpc, hash)
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
