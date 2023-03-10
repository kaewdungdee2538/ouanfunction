package ouanfunction

import (
	"github.com/kaewdungdee2538/ouanfunction/qrcode"
)

func ConvertStringQrCodeToByte(qrText string) []byte{
	return qrcode.GenerateQrCodeToBytes(qrText)
}