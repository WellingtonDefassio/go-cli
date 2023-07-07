package utils

import (
	cp "crypto/rand"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func GeneratePassword(size int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	if size < 8 {
		fmt.Println("size must be bigger then 7")
		return ""
	}
	b := make([]byte, size)
	if _, err := cp.Read(b); err != nil {
		fmt.Println("error to generated password")
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}
