package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// sha256: string to string
func Sha256Encode(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	md := hash.Sum(nil)
	encodedstr := hex.EncodeToString(md)
	return encodedstr
}
