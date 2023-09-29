package helper

import ("github.com/ayoseun/geth-lte/types")


func ToCallArg(msg types.ParamObject) interface{} {
arg := types.ParamObject{
	To:  msg.To,
	From: msg.From,
	Data: msg.Data,
}
return  []interface{}{arg, "latest"}
}
