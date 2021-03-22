package helps

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash256(data []byte) string {
	h := sha256.New()
	h.Write(data)
	var out []byte
	out = h.Sum(nil)

	return hex.EncodeToString(out)
}
