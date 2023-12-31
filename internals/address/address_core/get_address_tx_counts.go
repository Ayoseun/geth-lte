package address_core

import (
	"encoding/json"
	

	"github.com/ayoseun/geth-lte/common/hexutil"
	"github.com/ayoseun/geth-lte/rpc_calls"
	"github.com/ayoseun/geth-lte/types" // Import the JSONRPC package
)

func GetAddressTXCount(rpc string, address string) (string, error) {
	//"https://bsc.meowrpc.com"
	// Define the URL you want to send a POST request to
	url := rpc

	// Create a JSON-RPC request struct
	request := types.JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_getTransactionCount",
		Params:  []interface{}{address, "latest"},
		ID:      123,
	}

	// Specify the content type for the request
	contentType := "application/json"

	// Send the JSON-RPC request and handle the response
	response, err := rpc_calls.HttpRequest(url, request, contentType)
	if err != nil {
		return "", err
	}



	// Define a struct to represent the JSON response
	type JSONResponse struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  string `json:"result"`
	}

	// Create a variable to hold the JSON response
	var parsedResponse JSONResponse

	// Parse the JSON response into the struct
	err = json.Unmarshal([]byte(response), &parsedResponse)
	if err != nil {
		return "", err
	}

	result, err := hexutil.DecodeBig(parsedResponse.Result)
	if err != nil {
		return "", err
	}


	resultStr := result.String()


	return resultStr, nil
}
