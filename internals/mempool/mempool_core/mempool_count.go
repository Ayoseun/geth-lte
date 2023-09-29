package mempool_core

import (
	"encoding/json"
	 "github.com/ayoseun/geth-lte/common/hexutil"
	"github.com/ayoseun/geth-lte/rpc_calls"
	"github.com/ayoseun/geth-lte/types" // Import the JSONRPC package
)

func TxMemPoolCount(rpc string) (types.MemPoolCount, error) {
    // Define the URL you want to send a POST request to
    url := rpc

    // Create a JSON-RPC request struct
    request := types.JSONRPCRequest{
        JSONRPC: "2.0",
        Method:  "txpool_status",
        Params:  []interface{}{},
        ID:      123,
    }

    // Specify the content type for the request
    contentType := "application/json"

    // Send the JSON-RPC request and handle the response
    response, err := rpc_calls.HttpRequest(url, request, contentType)
    if err != nil {
	
    }

    // Define a struct to represent the JSON response
    var parsedResponse types.MemPoolCount



    // Parse the JSON response into the struct
    err = json.Unmarshal([]byte(response), &parsedResponse)
    if err != nil {
     
    }
   // Format the "Pending" and "Queued" values
   parsedResponse.Result.Pending,err = hexutil.FormatHex(parsedResponse.Result.Pending)
   parsedResponse.Result.Queued,err = hexutil.FormatHex(parsedResponse.Result.Queued)


    return parsedResponse, nil
}
