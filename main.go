package main

import (
	"fmt"

	"github.com/ayoseun/geth-lte/internals/address"
	"github.com/ayoseun/geth-lte/internals/block"
	"github.com/ayoseun/geth-lte/internals/contract"
	"github.com/ayoseun/geth-lte/internals/mempool"
)

//0x2c9387ec7a7c84c

func main() {
	ViewTxMemPool()
	
}

func userContractBalance() {
	// Define the RPC string
	rpc := "https://rpc.ankr.com/polygon"


	result, err := contract.BalanceOf(rpc,"0xbe991e317fbb4e82f7f99c62096f36b7a24278de","0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(result))
	
}



//Read MemPool Realtime
func ViewTxMemPool() {
	poolCh := make(chan string)
	go mempool.RTLTransactionsMempool("wss://polygon-mainnet.g.alchemy.com/v2/u2BAYrPLhxkHb7O9AE2mqr3snCXhDB63", poolCh)
	for value := range poolCh {
		fmt.Println("Real-time data:", value)
	}
}



// reading mempool
func MemPool() {
	result, err := mempool.TransactionsMemPool("https://bsc.meowrpc.com")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pool: %s\n", result)
}

// Get block hash
func BlockByHash() {
	result, err := block.GetBlockByHash("https://bsc.meowrpc.com", "0x499d2f7bcd2c37e869f6721edb690105d19275e2ae25911c7d81b75305075dcd")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n", result)
}

// Get Transaction hash
func GetTransactionByHash() {
	result, err := address.GetTransactionByHash("https://bsc.meowrpc.com", "0x8b2c166606569b4a4ed68a764a370f1dbc5a61e8d7a2ea4d1812cc7f9c487ee1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n", result)
}

// Get block transaction count
func GetBlockTransactionCounts() {
	result, err := address.GetAddressTransactionCount("https://bsc.meowrpc.com", "0x9e4cc336022fd3fdae0f5ad25b758f040f30040a73a45fca1be9e440bac91902")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n", result)
}

// latest block example
func LatestBlock() {
	result, err := block.GetLatestBlock("https://bsc.meowrpc.com")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Latest Block: %s\n", result)
}

// Get address transactions count
func AddressTransactionCount() {
	result, err := address.GetAddressTransactionCount("https://bsc.meowrpc.com", "0xa6f79B60359f141df90A0C745125B131cAAfFD12")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n", result)
}
