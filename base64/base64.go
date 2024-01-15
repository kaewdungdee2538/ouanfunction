package base64

import (
	"encoding/base64"
)

func ByteToBase64String(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Encode(plaintext string) string {
	return base64.StdEncoding.EncodeToString([]byte(plaintext))
}

func Decode(encrypt string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
