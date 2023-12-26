package loopring

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/loopring/go-loopring-sig/eddsa"
	"github.com/loopring/go-loopring-sig/poseidon"
	"github.com/loopring/go-loopring-sig/utils"
)

func GetLockHashAndEddsaSignature(
	privateKey string,
	exchangeAddress string,
	accountId string,
	tokenId string,
	volume string,
	timestamp string,
) (string, error) {
	privateKeyBig := new(big.Int)
	exchangeAddressBig := new(big.Int)
	accountIdBig := new(big.Int)
	tokenIdBig := new(big.Int)
	volumeBig := new(big.Int)
	timestampBig := new(big.Int)

	privateKeyBig.SetString(strings.TrimPrefix(privateKey, "0x"), 16)
	exchangeAddressBig.SetString(strings.TrimPrefix(exchangeAddress, "0x"), 16)
	accountIdBig.SetString(accountId, 10)
	tokenIdBig.SetString(tokenId, 10)
	volumeBig.SetString(volume, 10)
	timestampBig.SetString(timestamp, 10)

	hash, err := poseidon.HashWithParams([]*big.Int{
		exchangeAddressBig,
		accountIdBig,
		tokenIdBig,
		volumeBig,
		timestampBig,
	}, 53)
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
