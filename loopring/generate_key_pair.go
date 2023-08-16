package loopring

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/loopexchange-labs/go-loopring-sig/eddsa"
)

func bnToBufWithFixedLength(bn *big.Int, outputLength int) []int64 {
	hex := fmt.Sprintf("%x", bn)
	if len(hex)%2 > 0 {
		hex = fmt.Sprintf("0%s", hex)
	}
	length := len(hex) / 2

	u := make([]int64, length)
	var i, j int
	for i < length {
		u[i], _ = strconv.ParseInt(hex[j:j+2], 16, 64)
		i += 1
		j += 2
	}

	more := outputLength - len(u)
	if more > 0 {
		for i = 0; i < more; i++ {
			u = append([]int64{0}, u...)
		}
	} else {
		u = u[0:outputLength]
	}

	return u
}

func GenerateKeyPair(signature string) interface{} {
	signatureBigInt := new(big.Int)
	signatureBigInt.SetString(strings.TrimPrefix(signature, "0x"), 16)

	hash := sha256.Sum256(signatureBigInt.Bytes())
	seedBuff := hash[:]

	seed := new(big.Int)
	seed.SetBytes(seedBuff)

	bitIntDataItems := bnToBufWithFixedLength(seed, 32)

	sum := big.NewInt(0)
	for idx, item := range bitIntDataItems {
		n, t, i := big.NewInt(int64(item)), big.NewInt(256), big.NewInt(int64(idx))
		t = t.Exp(t, i, nil)
		n = n.Mul(n, t)
		sum = sum.Add(sum, n)
	}

	secretKey := sum.Mod(sum, eddsa.L)
	copySecretKey := new(big.Int).Set(secretKey)
	publicKey := eddsa.NewPoint().Mul(copySecretKey, eddsa.B8)

	return map[string]interface{}{
		"keyPair": map[string]interface{}{
			"publicKeyX": publicKey.X.String(),
			"publicKeyY": publicKey.Y.String(),
			"secretKey":  secretKey.String(),
		},
		"formatedPx": fmt.Sprintf("0x%064s", publicKey.X.Text(16)),
		"formatedPy": fmt.Sprintf("0x%064s", publicKey.Y.Text(16)),
		"sk":         fmt.Sprintf("0x%s", secretKey.Text(16)),
	}
}
