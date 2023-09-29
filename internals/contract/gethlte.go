package contract

import (
	"encoding/json"


	"github.com/ayoseun/geth-lte/types"

	"github.com/ayoseun/geth-lte/internals/contract/contract_core"
	"github.com/ayoseun/geth-lte/types/results"
)


func BalanceOf(rpc,address string,contractAddress string) ([]byte,error) {
	// Define the RPC string
	

	// Create an instance of CallMsg and fill it with data
	msg := types.ParamObject{
		To:   contractAddress,
		From: address,
		Data: "0x70a08231000000000000000000000000be991e317fbb4e82f7f99c62096f36b7a24278de",
	
	}
	// Call the BalanceOf function
	balance, err := contract_core.BalanceOf(rpc, msg)
	if err != nil {
		return nil, err
	}

	contractData := results.AddressContractBalanceResponse{
        Address: address,
		Balance:  balance.Result,
	}
	contractJSON, err := json.Marshal(contractData)
	if err != nil {

		// Return an error if there's a problem fetching the balance
		return contractJSON, err
	}

	// Return the JSON-encoded response as a []byte

	return contractJSON,nil
	
}