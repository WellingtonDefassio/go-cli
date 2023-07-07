package utils

import "encoding/base64"

func EncodedString(s string) string {
	encode := base64.StdEncoding.EncodeToString([]byte(s))
	return encode
}

func DecodeString(s string) string {
	decodeString, _ := base64.StdEncoding.DecodeString(s)
	return string(decodeString)
}
