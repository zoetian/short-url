package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/mr-tron/base58"
)

func sha2560f(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoded := base58.Encode(bytes)
	return string(encoded)
}

func GenerateShortUrl(longUrl, userID string) string {
	urlHashBytes := sha2560f(longUrl + userID)
	num := new(big.Int).SetBytes(urlHashBytes).Uint64()
	res := base58Encoded([]byte(fmt.Sprintf("%d", num)))
	return res[:8]
}
