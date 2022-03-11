package general

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func GetMD5Hash(text string, text_2 string) string {
	hasher := md5.New()
	hasher.Write([]byte(text + text_2))
	return hex.EncodeToString(hasher.Sum(nil))
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	random_string := make([]rune, length)
	for index := range random_string {
		random_string[index] = letters[rand.Intn(len(letters))]
	}
	return string(random_string)
}
