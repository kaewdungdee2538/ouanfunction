package ksherPayment

import (
	"github.com/ksher-api/ksher-sdk/go/src/KsherGO"
)

func GenerateSign(privateKeyPath string) (string, error) {
	privateKeyData := KsherGO.ReadPrivateKeyFromPath(privateKeyPath)
	return KsherGO.KsherSign(nil, privateKeyData)
}
