package transition

import (
	"errors"
)

// GetValue 转换 input 中的字段
func GetValue(value string) (string, error) {
	m := make(map[string]string)
	m["Transfer"] = "0x"
	m["Multicall"] = "0x5ae401dc"
	m["Approve"] = "0x095ea7b3"
	m["AddLiquidityETH"] = "0xf305d719"
	m["EnableTrading"] = "0x8a8c523c"
	m["renounceOwnership"] = "0x715018a6"
	m["LockLPToken"] = "0x8af416f6"
	m["lockTokens"] = "0x7d533c1e"
	m["disableTransferDelay"] = "0xe884f260"
	m["updateMaxWalletAmount"] = "0xc18bc195"
	m["removeLiquidityWithPermit"] = "0x2195995c"
	m["updateMaxTxnAmount"] = "0x203e727e"
	m["updateBuyFees"] = "0x8095d564"
	m["updateSellFees"] = "0xc17b5b8c"
	m["transferLockOwnership"] = "0xbef497fd"
	m["swapExactTokensForETH"] = "0x18cbafe5"
	m["relock"] = "0x60491d24"

	for k, v := range m {
		if value == v {
			return k, nil
		}
	}

	return "", errors.New("not found")
}
