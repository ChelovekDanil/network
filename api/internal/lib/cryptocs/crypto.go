package cryptocs

import (
	"crypto/sha512"
	"encoding/hex"
)

func Hash(data string) string {
	hash := sha512.New()

	hash.Write([]byte(data))

	hashedData := hash.Sum(nil)

	hashedString := hex.EncodeToString(hashedData)

	return hashedString
}
