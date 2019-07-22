package encode

import (
	"github.com/btcsuite/btcutil/base58"
)

func EncodeB58(data []byte) string {
	return base58.Encode(data)
}
