/*
	copyright © 2022 by @https://github.com/Chuiko-GIT
*/

package binance

// Block - define crypto type.
const (
	SearchCryptoBTC = "BTC"
	SearchCryptoETH = "ETH"
	SearchCryptoXLM = "XLM"

	CryptoTypeUSDT = "USDT"
)

// CryptoToUSDT - handle crypto to USDT command.
var CryptoToUSDT = map[SearchCrypto]struct{}{
	SearchCryptoBTC: {},
	SearchCryptoETH: {},
	SearchCryptoXLM: {},
}

type (
	// SearchCrypto - define crypto type.
	SearchCrypto string

	// binanceResponseCourseToUSDT - define binance response course to USDT.
	binanceResponseCourseToUSDT struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}
)

// Validate - validate crypto to USDT.
func (s SearchCrypto) Validate() bool {
	_, ok := CryptoToUSDT[s]
	return ok
}
