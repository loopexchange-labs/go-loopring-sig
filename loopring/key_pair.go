package loopring

import (
	"fmt"
	"math/big"
)

type KeyPairFormatted struct {
	SecretKey  string
	PublicKeyX string
	PublicKeyY string
}

type KeyPair struct {
	SecretKey  *big.Int
	PublicKeyX *big.Int
	PublicKeyY *big.Int
}

func (kp KeyPair) FormatSecretKey() string {
	return fmt.Sprintf("0x%s", kp.SecretKey.Text(16))
}

func (kp KeyPair) FormatPublicKeyX() string {
	return fmt.Sprintf("0x%064s", kp.PublicKeyX.Text(16))
}

func (kp KeyPair) FormatPublicKeyY() string {
	return fmt.Sprintf("0x%064s", kp.PublicKeyY.Text(16))
}

func (kp KeyPair) ToFormatted() KeyPairFormatted {
	return KeyPairFormatted{
		SecretKey:  fmt.Sprintf("0x%s", kp.SecretKey.Text(16)),
		PublicKeyX: fmt.Sprintf("0x%064s", kp.PublicKeyX.Text(16)),
		PublicKeyY: fmt.Sprintf("0x%064s", kp.PublicKeyY.Text(16)),
	}
}
