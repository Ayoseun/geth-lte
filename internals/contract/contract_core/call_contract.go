package contract_core

import (
	"encoding/json"
	"fmt"

	"github.com/ayoseun/geth-lte/internals/contract/helper"
	"github.com/ayoseun/geth-lte/rpc_calls"
	"github.com/ayoseun/geth-lte/types" // Import the JSONRPC package
	"github.com/ayoseun/geth-lte/utils"
)

func BalanceOf(rpc string, msg types.ParamObject) (types.JSONRPCResult, error) {
	// Define the URL you want to send a POST request to
	url := rpc
	fmt.Println(msg)

	// Create a JSON-RPC request struct
	request := types.JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_call",
		Params:  helper.ToCallArg(msg),

		ID: 123,
	}

	// Specify the content type for the request
	contentType := "application/json"

	// Send the JSON-RPC request and handle the response
	response, err := rpc_calls.HttpRequest(url, request, contentType)
	if err != nil {

	}

	// Define a struct to represent the JSON response
	var parsedResponse types.JSONRPCResult

	// Parse the JSON response into the struct
	err = json.Unmarshal([]byte(response), &parsedResponse)
	if err != nil {

	}

	parsedResponse.Result, err = utils.HexToString(parsedResponse.Result)

	return parsedResponse, nil
}
