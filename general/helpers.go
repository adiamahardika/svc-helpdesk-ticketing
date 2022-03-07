package general

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(text string, text_2 string) string {
	hasher := md5.New()
	hasher.Write([]byte(text + text_2))
	return hex.EncodeToString(hasher.Sum(nil))
}
