package internals

import (
	"encoding/json"
	"github.com/ayoseun/geth-lte/common"
	"github.com/ayoseun/geth-lte/rpc_calls"
	"github.com/ayoseun/geth-lte/types" // Import the JSONRPC package
)

func GetBalance(rpc string, address string) (string, error) {
	//"https://bsc.meowrpc.com"
	// Define the URL you want to send a POST request to
	url := rpc

	// Create a JSON-RPC request struct
	request := types.JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_getBalance",
		Params:  []interface{}{address, "latest"},
		ID:      123,
	}

	// Specify the content type for the request
	contentType := "application/json"

	// Send the JSON-RPC request and handle the response
	response, err := rpc_calls.GetBalanceRequest(url, request, contentType)
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

	denominatorStr := "1000000000000000000"
	// precision := 2
	resultStr := result.String()
	result6, err := hexutil.DivideLargeNumbers(resultStr, denominatorStr)
	if err != nil {
		return "", err
	}

	return result6, nil
}
