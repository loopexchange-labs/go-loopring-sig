package loopring

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/loopring/go-loopring-sig/eddsa"
	"github.com/loopring/go-loopring-sig/poseidon"
	"github.com/loopring/go-loopring-sig/utils"
)

func GetEddsaSigNftOrder(
	privateKey string,
	exchangeAddress string,
	storageId string,
	accountId string,
	sellTokenId string,
	buyTokenId string,
	sellTokenAmount string,
	buyTokenAmount string,
	validUntil string,
	maxFeeBips string,
	fillAmountBOrS string,
	takerAddress string,
) (string, error) {
	privateKeyBig := new(big.Int)
	exchangeAddressBig := new(big.Int)
	storageIdBig := new(big.Int)
	accountIdBig := new(big.Int)
	sellTokenIdBig := new(big.Int)
	buyTokenIdBig := new(big.Int)
	sellTokenAmountBig := new(big.Int)
	buyTokenAmountBig := new(big.Int)
	validUntilBig := new(big.Int)
	maxFeeBipsBig := new(big.Int)
	fillAmountBOrSBig := new(big.Int)
	takerAddressBig := new(big.Int)

	privateKeyBig.SetString(strings.TrimPrefix(privateKey, "0x"), 16)
	exchangeAddressBig.SetString(strings.TrimPrefix(exchangeAddress, "0x"), 16)
	storageIdBig.SetString(storageId, 10)
	accountIdBig.SetString(accountId, 10)
	sellTokenIdBig.SetString(sellTokenId, 10)
	if strings.HasPrefix(buyTokenId, "0x") {
		buyTokenIdBig.SetString(strings.TrimPrefix(buyTokenId, "0x"), 16)
	} else {
		buyTokenIdBig.SetString(buyTokenId, 10)
	}
	sellTokenAmountBig.SetString(sellTokenAmount, 10)
	buyTokenAmountBig.SetString(buyTokenAmount, 10)
	validUntilBig.SetString(validUntil, 10)
	maxFeeBipsBig.SetString(maxFeeBips, 10)
	fillAmountBOrSBig.SetString(fillAmountBOrS, 10)
	takerAddressBig.SetString(strings.TrimPrefix(takerAddress, "0x"), 16)

	hash, err := poseidon.Hash([]*big.Int{
		exchangeAddressBig,
		storageIdBig,
		accountIdBig,
		sellTokenIdBig,
		buyTokenIdBig,
		sellTokenAmountBig,
		buyTokenAmountBig,
		validUntilBig,
		maxFeeBipsBig,
		fillAmountBOrSBig,
		takerAddressBig,
	})
	if err != nil {
		return "", err
	}

	var pk eddsa.PrivateKey //lint:ignore S1021 conversion
	pk = utils.BigIntLEBytes(privateKeyBig)
	sig := pk.SignPoseidon(hash)

	return "0x" +
		fmt.Sprintf("%064s", sig.R8.X.Text(16)) +
		fmt.Sprintf("%064s", sig.R8.Y.Text(16)) +
		fmt.Sprintf("%064s", sig.S.Text(16)), nil
}
