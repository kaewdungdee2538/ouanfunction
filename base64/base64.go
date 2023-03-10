package base64

import (
	"encoding/base64"
)

func ByteToBase64String(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}