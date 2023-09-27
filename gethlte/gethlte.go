package gethlte

import (
	"encoding/json"

	"github.com/ayoseun/geth-lte/internals"
	"github.com/ayoseun/geth-lte/internals/address"
	"github.com/ayoseun/geth-lte/internals/block"
	"github.com/ayoseun/geth-lte/types"
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
    tx, err := address.GetTransactionByHash(rpc, hash)
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