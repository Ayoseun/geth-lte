package main

import (
	"fmt"

	"github.com/ayoseun/geth-lte/gethlte"
	"github.com/ayoseun/geth-lte/utils"
)

//0x2c9387ec7a7c84c

func main(){
	result,err:= gethlte.GetBlockByHash("https://bsc.meowrpc.com","0x499d2f7bcd2c37e869f6721edb690105d19275e2ae25911c7d81b75305075dcd")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n",result)
}



//Get Transaction hash
func GetTransactionByHash(){
	result,err:= gethlte.GetTransactionByHash("https://bsc.meowrpc.com","0x8b2c166606569b4a4ed68a764a370f1dbc5a61e8d7a2ea4d1812cc7f9c487ee1")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n",result)
}




//Get block transaction count
func GetBlockTransactionCounts(){
	result,err:= gethlte.GetBlockTransactionCount("https://bsc.meowrpc.com","0x9e4cc336022fd3fdae0f5ad25b758f040f30040a73a45fca1be9e440bac91902")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n",result)
}


//latest block example
func LatestBlock(){
	result,err:= gethlte.GetLatestBlock("https://bsc.meowrpc.com")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Latest Block: %s\n",result)
}




//Get address transactions count
func AddressTransactionCount(){
	result,err:= gethlte.GetAddressTransactionCount("https://bsc.meowrpc.com","0xa6f79B60359f141df90A0C745125B131cAAfFD12")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n",result)
}

















//get gas price 

func getgasprice(){
	result,err:= gethlte.GetGasPrice("https://bsc.meowrpc.com")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n",result)
}




//Get wallet balance example
func showBalane(){
	result,err:= gethlte.GetBalance("https://bsc.meowrpc.com","0xa6f79B60359f141df90A0C745125B131cAAfFD12")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ether balance: %s\n",result)
}

// format ethers example
func convertHex2Ether(){
	result,err:= utils.FormatEther("0x2c9387ec7a7c84c")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Eth: %s\n",result)
}


//parse ether example
func parseEther(){
	result,err:= utils.ParseEther("2.0",16)
	if err != nil {
	panic(err)
	}
	fmt.Printf("Eth: %s\n",result)
}