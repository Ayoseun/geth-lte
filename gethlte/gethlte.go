package gethlte

import (
	"encoding/json"
	"github.com/ayoseun/geth-lte/internals"
	"github.com/ayoseun/geth-lte/types"
)




func GetBalance(rpc string,walletAddress string)([]byte, error) {
//"https://bsc.meowrpc.com"
	walletBalance, err := internals.GetBalance(rpc, walletAddress)


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




