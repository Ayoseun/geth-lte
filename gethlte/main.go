package gethlte

import (
	"encoding/json"
	"fmt"
	"github.com/ayoseun/geth-lte/internals"
	"github.com/ayoseun/geth-lte/types"
)



//help
func GetBalance(rpc string,walletAddress string) {
//"https://bsc.meowrpc.com"
	walletBalance, err := internals.GetBalance(rpc, walletAddress)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	response := types.WalletBalanceResponse{
		Balance: walletBalance,
		Address: walletAddress,
	}

	// Marshal the response to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(responseJSON))


}


