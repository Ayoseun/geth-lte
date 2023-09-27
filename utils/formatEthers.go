package utils

import (


	hexutil "github.com/ayoseun/geth-lte/common"
)


func FormatEther(hex string,precision ...int)(string, error) {

	result, err := hexutil.DecodeBig(hex)
	if err != nil {
		return "", err
	}

	denominatorStr := "1000000000000000000"

	resultStr := result.String()
		// Default precision to 8 if not provided
		if len(precision) == 0 {
			precision = append(precision, 18)
		}
	etherStr, err := hexutil.DivideLargeNumbers(resultStr, denominatorStr,precision[0])
	if err != nil {
		return "", err
	}

	return etherStr, nil
}


func ParseEther(big string,precision ...int)(string, error) {




	// Default precision to 8 if not provided
	if len(precision) == 0 {
		precision = append(precision, 18)
	}
	
	etherStr, err := hexutil.MultiplyDecimal(big,precision[0])
	if err != nil {
		return "", err
	}

	return etherStr, nil
}