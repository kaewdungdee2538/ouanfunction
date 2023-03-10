package qrcode

import qrcode "github.com/skip2/go-qrcode"

func GenerateQrCodeToBytes(qrText string) ([]byte){
	if png, err := qrcode.Encode(qrText, qrcode.Medium, 256);err==nil{
		return png
	}else{
		return nil
	}
}
