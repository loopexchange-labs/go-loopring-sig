package loopring

import (
	"math/big"
	"strings"

	"github.com/loopexchange-labs/go-loopring-sig/eddsa"
	"github.com/loopexchange-labs/go-loopring-sig/utils"
)

func NewPrivateKeyFromString(privateKeyString string) *eddsa.PrivateKey {
	privateKeyBig := new(big.Int)
	privateKeyBig.SetString(strings.TrimPrefix(privateKeyString, "0x"), 16)

	var pk eddsa.PrivateKey //lint:ignore S1021 conversion
	pk = utils.BigIntLEBytes(privateKeyBig)

	return &pk
}
